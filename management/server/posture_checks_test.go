package server

import (
	"context"
	"testing"
	"time"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/netbirdio/netbird/management/server/group"

	"github.com/netbirdio/netbird/management/server/posture"
)

const (
	adminUserID      = "adminUserID"
	regularUserID    = "regularUserID"
	postureCheckName = "Existing check"
)

func TestDefaultAccountManager_PostureCheck(t *testing.T) {
	am, err := createManager(t)
	if err != nil {
		t.Error("failed to create account manager")
	}

	account, err := initTestPostureChecksAccount(am)
	if err != nil {
		t.Error("failed to init testing account")
	}

	t.Run("Generic posture check flow", func(t *testing.T) {
		// regular users can not create checks
		_, err = am.SavePostureChecks(context.Background(), account.Id, regularUserID, &posture.Checks{})
		assert.Error(t, err)

		// regular users cannot list check
		_, err = am.ListPostureChecks(context.Background(), account.Id, regularUserID)
		assert.Error(t, err)

		// should be possible to create posture check with uniq name
		postureCheck, err := am.SavePostureChecks(context.Background(), account.Id, adminUserID, &posture.Checks{
			Name: postureCheckName,
			Checks: posture.ChecksDefinition{
				NBVersionCheck: &posture.NBVersionCheck{
					MinVersion: "0.26.0",
				},
			},
		})
		assert.NoError(t, err)

		// admin users can list check
		checks, err := am.ListPostureChecks(context.Background(), account.Id, adminUserID)
		assert.NoError(t, err)
		assert.Len(t, checks, 1)

		// should not be possible to create posture check with non uniq name
		_, err = am.SavePostureChecks(context.Background(), account.Id, adminUserID, &posture.Checks{
			Name: postureCheckName,
			Checks: posture.ChecksDefinition{
				GeoLocationCheck: &posture.GeoLocationCheck{
					Locations: []posture.Location{
						{
							CountryCode: "DE",
						},
					},
				},
			},
		})
		assert.Error(t, err)

		// admins can update posture checks
		postureCheck.Checks = posture.ChecksDefinition{
			NBVersionCheck: &posture.NBVersionCheck{
				MinVersion: "0.27.0",
			},
		}
		_, err = am.SavePostureChecks(context.Background(), account.Id, adminUserID, postureCheck)
		assert.NoError(t, err)

		// users should not be able to delete posture checks
		err = am.DeletePostureChecks(context.Background(), account.Id, postureCheck.ID, regularUserID)
		assert.Error(t, err)

		// admin should be able to delete posture checks
		err = am.DeletePostureChecks(context.Background(), account.Id, postureCheck.ID, adminUserID)
		assert.NoError(t, err)
		checks, err = am.ListPostureChecks(context.Background(), account.Id, adminUserID)
		assert.NoError(t, err)
		assert.Len(t, checks, 0)
	})
}

func initTestPostureChecksAccount(am *DefaultAccountManager) (*Account, error) {
	accountID := "testingAccount"
	domain := "example.com"

	admin := &User{
		Id:   adminUserID,
		Role: UserRoleAdmin,
	}
	user := &User{
		Id:   regularUserID,
		Role: UserRoleUser,
	}

	account := newAccountWithId(context.Background(), accountID, groupAdminUserID, domain)
	account.Users[admin.Id] = admin
	account.Users[user.Id] = user

	err := am.Store.SaveAccount(context.Background(), account)
	if err != nil {
		return nil, err
	}

	return am.Store.GetAccount(context.Background(), account.Id)
}

func TestPostureCheckAccountPeersUpdate(t *testing.T) {
	manager, account, peer1, peer2, peer3 := setupNetworkMapTest(t)

	err := manager.SaveGroups(context.Background(), account.Id, userID, []*group.Group{
		{
			ID:    "groupA",
			Name:  "GroupA",
			Peers: []string{peer1.ID, peer2.ID, peer3.ID},
		},
		{
			ID:    "groupB",
			Name:  "GroupB",
			Peers: []string{},
		},
		{
			ID:    "groupC",
			Name:  "GroupC",
			Peers: []string{},
		},
	})
	assert.NoError(t, err)

	updMsg := manager.peersUpdateManager.CreateChannel(context.Background(), peer1.ID)
	t.Cleanup(func() {
		manager.peersUpdateManager.CloseChannel(context.Background(), peer1.ID)
	})

	postureCheckA := &posture.Checks{
		Name:      "postureCheckA",
		AccountID: account.Id,
		Checks: posture.ChecksDefinition{
			ProcessCheck: &posture.ProcessCheck{
				Processes: []posture.Process{
					{LinuxPath: "/usr/bin/netbird", MacPath: "/usr/local/bin/netbird"},
				},
			},
		},
	}
	postureCheckA, err = manager.SavePostureChecks(context.Background(), account.Id, userID, postureCheckA)
	require.NoError(t, err)

	postureCheckB := &posture.Checks{
		Name:      "postureCheckB",
		AccountID: account.Id,
		Checks: posture.ChecksDefinition{
			NBVersionCheck: &posture.NBVersionCheck{
				MinVersion: "0.28.0",
			},
		},
	}

	// Saving unused posture check should not update account peers and not send peer update
	t.Run("saving unused posture check", func(t *testing.T) {
		done := make(chan struct{})
		go func() {
			peerShouldNotReceiveUpdate(t, updMsg)
			close(done)
		}()

		postureCheckB, err = manager.SavePostureChecks(context.Background(), account.Id, userID, postureCheckB)
		assert.NoError(t, err)

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("timeout waiting for peerShouldNotReceiveUpdate")
		}
	})

	// Updating unused posture check should not update account peers and not send peer update
	t.Run("updating unused posture check", func(t *testing.T) {
		done := make(chan struct{})
		go func() {
			peerShouldNotReceiveUpdate(t, updMsg)
			close(done)
		}()

		postureCheckB.Checks = posture.ChecksDefinition{
			NBVersionCheck: &posture.NBVersionCheck{
				MinVersion: "0.29.0",
			},
		}
		_, err = manager.SavePostureChecks(context.Background(), account.Id, userID, postureCheckB)
		assert.NoError(t, err)

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("timeout waiting for peerShouldNotReceiveUpdate")
		}
	})

	policy := Policy{
		ID:      "policyA",
		Enabled: true,
		Rules: []*PolicyRule{
			{
				ID:            xid.New().String(),
				Enabled:       true,
				Sources:       []string{"groupA"},
				Destinations:  []string{"groupA"},
				Bidirectional: true,
				Action:        PolicyTrafficActionAccept,
			},
		},
		SourcePostureChecks: []string{postureCheckB.ID},
	}

	// Linking posture check to policy should trigger update account peers and send peer update
	t.Run("linking posture check to policy with peers", func(t *testing.T) {
		done := make(chan struct{})
		go func() {
			peerShouldReceiveUpdate(t, updMsg)
			close(done)
		}()

		err := manager.SavePolicy(context.Background(), account.Id, userID, &policy, false)
		assert.NoError(t, err)

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("timeout waiting for peerShouldReceiveUpdate")
		}
	})

	// Updating linked posture checks should update account peers and send peer update
	t.Run("updating linked to posture check with peers", func(t *testing.T) {
		postureCheckB.Checks = posture.ChecksDefinition{
			NBVersionCheck: &posture.NBVersionCheck{
				MinVersion: "0.29.0",
			},
			ProcessCheck: &posture.ProcessCheck{
				Processes: []posture.Process{
					{LinuxPath: "/usr/bin/netbird", MacPath: "/usr/local/bin/netbird"},
				},
			},
		}

		done := make(chan struct{})
		go func() {
			peerShouldReceiveUpdate(t, updMsg)
			close(done)
		}()

		_, err = manager.SavePostureChecks(context.Background(), account.Id, userID, postureCheckB)
		assert.NoError(t, err)

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("timeout waiting for peerShouldReceiveUpdate")
		}
	})

	// Removing posture check from policy should trigger account peers update and send peer update
	t.Run("removing posture check from policy", func(t *testing.T) {
		done := make(chan struct{})
		go func() {
			peerShouldReceiveUpdate(t, updMsg)
			close(done)
		}()

		policy.SourcePostureChecks = []string{}

		err := manager.SavePolicy(context.Background(), account.Id, userID, &policy, true)
		assert.NoError(t, err)

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("timeout waiting for peerShouldReceiveUpdate")
		}
	})

	// Deleting unused posture check should not trigger account peers update and not send peer update
	t.Run("deleting unused posture check", func(t *testing.T) {
		done := make(chan struct{})
		go func() {
			peerShouldNotReceiveUpdate(t, updMsg)
			close(done)
		}()

		err := manager.DeletePostureChecks(context.Background(), account.Id, postureCheckA.ID, userID)
		assert.NoError(t, err)

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("timeout waiting for peerShouldNotReceiveUpdate")
		}
	})

	_, err = manager.SavePostureChecks(context.Background(), account.Id, userID, postureCheckB)
	assert.NoError(t, err)

	// Updating linked posture check to policy with no peers should not trigger account peers update and not send peer update
	t.Run("updating linked posture check to policy with no peers", func(t *testing.T) {
		policy = Policy{
			ID:      "policyB",
			Enabled: true,
			Rules: []*PolicyRule{
				{
					ID:            xid.New().String(),
					Enabled:       true,
					Sources:       []string{"groupB"},
					Destinations:  []string{"groupC"},
					Bidirectional: true,
					Action:        PolicyTrafficActionAccept,
				},
			},
			SourcePostureChecks: []string{postureCheckB.ID},
		}
		err = manager.SavePolicy(context.Background(), account.Id, userID, &policy, false)
		assert.NoError(t, err)

		done := make(chan struct{})
		go func() {
			peerShouldNotReceiveUpdate(t, updMsg)
			close(done)
		}()

		postureCheckB.Checks = posture.ChecksDefinition{
			NBVersionCheck: &posture.NBVersionCheck{
				MinVersion: "0.29.0",
			},
		}
		_, err = manager.SavePostureChecks(context.Background(), account.Id, userID, postureCheckB)
		assert.NoError(t, err)

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("timeout waiting for peerShouldNotReceiveUpdate")
		}
	})

	// Updating linked posture check to policy where destination has peers but source does not
	// should trigger account peers update and send peer update
	t.Run("updating linked posture check to policy where destination has peers but source does not", func(t *testing.T) {
		updMsg1 := manager.peersUpdateManager.CreateChannel(context.Background(), peer2.ID)
		t.Cleanup(func() {
			manager.peersUpdateManager.CloseChannel(context.Background(), peer2.ID)
		})
		policy = Policy{
			ID:      "policyB",
			Enabled: true,
			Rules: []*PolicyRule{
				{
					ID:            xid.New().String(),
					Enabled:       true,
					Sources:       []string{"groupB"},
					Destinations:  []string{"groupA"},
					Bidirectional: true,
					Action:        PolicyTrafficActionAccept,
				},
			},
			SourcePostureChecks: []string{postureCheckB.ID},
		}

		err = manager.SavePolicy(context.Background(), account.Id, userID, &policy, true)
		assert.NoError(t, err)

		done := make(chan struct{})
		go func() {
			peerShouldReceiveUpdate(t, updMsg1)
			close(done)
		}()

		postureCheckB.Checks = posture.ChecksDefinition{
			NBVersionCheck: &posture.NBVersionCheck{
				MinVersion: "0.29.0",
			},
		}
		_, err = manager.SavePostureChecks(context.Background(), account.Id, userID, postureCheckB)
		assert.NoError(t, err)

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("timeout waiting for peerShouldReceiveUpdate")
		}
	})

	// Updating linked client posture check to policy where source has peers but destination does not,
	// should trigger account peers update and send peer update
	t.Run("updating linked posture check to policy where source has peers but destination does not", func(t *testing.T) {
		policy = Policy{
			ID:      "policyB",
			Enabled: true,
			Rules: []*PolicyRule{
				{
					Enabled:       true,
					Sources:       []string{"groupA"},
					Destinations:  []string{"groupB"},
					Bidirectional: true,
					Action:        PolicyTrafficActionAccept,
				},
			},
			SourcePostureChecks: []string{postureCheckB.ID},
		}
		err = manager.SavePolicy(context.Background(), account.Id, userID, &policy, true)
		assert.NoError(t, err)

		done := make(chan struct{})
		go func() {
			peerShouldReceiveUpdate(t, updMsg)
			close(done)
		}()

		postureCheckB.Checks = posture.ChecksDefinition{
			ProcessCheck: &posture.ProcessCheck{
				Processes: []posture.Process{
					{
						LinuxPath: "/usr/bin/netbird",
					},
				},
			},
		}
		_, err = manager.SavePostureChecks(context.Background(), account.Id, userID, postureCheckB)
		assert.NoError(t, err)

		select {
		case <-done:
		case <-time.After(time.Second):
			t.Error("timeout waiting for peerShouldReceiveUpdate")
		}
	})
}

func TestArePostureCheckChangesAffectPeers(t *testing.T) {
	manager, err := createManager(t)
	require.NoError(t, err, "failed to create account manager")

	account, err := initTestPostureChecksAccount(manager)
	require.NoError(t, err, "failed to init testing account")

	groupA := &group.Group{
		ID:        "groupA",
		AccountID: account.Id,
		Peers:     []string{"peer1"},
	}

	groupB := &group.Group{
		ID:        "groupB",
		AccountID: account.Id,
		Peers:     []string{},
	}
	err = manager.Store.SaveGroups(context.Background(), LockingStrengthUpdate, []*group.Group{groupA, groupB})
	require.NoError(t, err, "failed to save groups")

	postureCheckA := &posture.Checks{
		Name:      "checkA",
		AccountID: account.Id,
		Checks: posture.ChecksDefinition{
			NBVersionCheck: &posture.NBVersionCheck{MinVersion: "0.33.1"},
		},
	}
	postureCheckA, err = manager.SavePostureChecks(context.Background(), account.Id, adminUserID, postureCheckA)
	require.NoError(t, err, "failed to save postureCheckA")

	postureCheckB := &posture.Checks{
		Name:      "checkB",
		AccountID: account.Id,
		Checks: posture.ChecksDefinition{
			NBVersionCheck: &posture.NBVersionCheck{MinVersion: "0.33.1"},
		},
	}
	postureCheckB, err = manager.SavePostureChecks(context.Background(), account.Id, adminUserID, postureCheckB)
	require.NoError(t, err, "failed to save postureCheckB")

	policy := &Policy{
		ID:        "policyA",
		AccountID: account.Id,
		Rules: []*PolicyRule{
			{
				ID:           "ruleA",
				PolicyID:     "policyA",
				Enabled:      true,
				Sources:      []string{"groupA"},
				Destinations: []string{"groupA"},
			},
		},
		SourcePostureChecks: []string{postureCheckA.ID},
	}

	err = manager.SavePolicy(context.Background(), account.Id, userID, policy, false)
	require.NoError(t, err, "failed to save policy")

	t.Run("posture check exists and is linked to policy with peers", func(t *testing.T) {
		result, err := arePostureCheckChangesAffectPeers(context.Background(), manager.Store, account.Id, postureCheckA.ID)
		require.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("posture check exists but is not linked to any policy", func(t *testing.T) {
		result, err := arePostureCheckChangesAffectPeers(context.Background(), manager.Store, account.Id, postureCheckB.ID)
		require.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("posture check does not exist", func(t *testing.T) {
		result, err := arePostureCheckChangesAffectPeers(context.Background(), manager.Store, account.Id, "unknown")
		require.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("posture check is linked to policy with no peers in source groups", func(t *testing.T) {
		policy.Rules[0].Sources = []string{"groupB"}
		policy.Rules[0].Destinations = []string{"groupA"}
		err = manager.SavePolicy(context.Background(), account.Id, userID, policy, true)
		require.NoError(t, err, "failed to update policy")

		result, err := arePostureCheckChangesAffectPeers(context.Background(), manager.Store, account.Id, postureCheckA.ID)
		require.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("posture check is linked to policy with no peers in destination groups", func(t *testing.T) {
		policy.Rules[0].Sources = []string{"groupA"}
		policy.Rules[0].Destinations = []string{"groupB"}
		err = manager.SavePolicy(context.Background(), account.Id, userID, policy, true)
		require.NoError(t, err, "failed to update policy")

		result, err := arePostureCheckChangesAffectPeers(context.Background(), manager.Store, account.Id, postureCheckA.ID)
		require.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("posture check is linked to policy but no peers in groups", func(t *testing.T) {
		groupA.Peers = []string{}
		err = manager.Store.SaveGroup(context.Background(), LockingStrengthUpdate, groupA)
		require.NoError(t, err, "failed to save groups")

		result, err := arePostureCheckChangesAffectPeers(context.Background(), manager.Store, account.Id, postureCheckA.ID)
		require.NoError(t, err)
		assert.False(t, result)
	})

	t.Run("posture check is linked to policy with non-existent group", func(t *testing.T) {
		policy.Rules[0].Sources = []string{"nonExistentGroup"}
		policy.Rules[0].Destinations = []string{"nonExistentGroup"}
		err = manager.SavePolicy(context.Background(), account.Id, userID, policy, true)
		require.NoError(t, err, "failed to update policy")

		result, err := arePostureCheckChangesAffectPeers(context.Background(), manager.Store, account.Id, postureCheckA.ID)
		require.NoError(t, err)
		assert.False(t, result)
	})
}
