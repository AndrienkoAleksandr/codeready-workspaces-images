//
// Copyright (c) 2021 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
//

package dashboard

import (
	"fmt"

	rbacv1 "k8s.io/api/rbac/v1"
)

const ClusterPermissionsDashboardFinalizer = "dashboard.clusterpermissions.finalizers.che.eclipse.org"

const DashboardSA = "che-dashboard"
const DashboardSAClusterRoleTemplate = "%s-che-dashboard"
const DashboardSAClusterRoleBindingTemplate = "%s-che-dashboard"

func GetPrivilegedPoliciesRulesForKubernetes() []rbacv1.PolicyRule {
	return []rbacv1.PolicyRule{
		{
			APIGroups: []string{"workspace.devfile.io"},
			Resources: []string{"devworkspaces"},
			Verbs:     []string{"create", "update", "patch", "get", "watch", "list", "delete"},
		},
		{
			APIGroups: []string{"workspace.devfile.io"},
			Resources: []string{"devworkspacetemplates"},
			Verbs:     []string{"create", "get", "list", "update", "patch", "delete"},
		},
		{
			APIGroups: []string{""},
			Resources: []string{"namespaces"},
			Verbs:     []string{"get", "create", "update", "list"},
		},
	}
}

func (d *Dashboard) getClusterRoleName() string {
	return fmt.Sprintf(DashboardSAClusterRoleTemplate, d.deployContext.CheCluster.Namespace)
}

func (d *Dashboard) getClusterRoleBindingName() string {
	return fmt.Sprintf(DashboardSAClusterRoleBindingTemplate, d.deployContext.CheCluster.Namespace)
}
