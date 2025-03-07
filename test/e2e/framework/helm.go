/*
Copyright 2021 RadonDB.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	. "github.com/onsi/gomega"
)

func HelmInstallChart(release, ns string) {
	strs := strings.Split(TestContext.OperatorImagePath, ":")
	if len(strs) != 2 {
		Failf(fmt.Sprintf("Invalid operator image path: %s", TestContext.OperatorImagePath))
	}
	args := []string{
		"install", release, "./" + TestContext.ChartPath,
		"--namespace", ns,
		"--values", TestContext.ChartValues, "--wait",
		"--kube-context", TestContext.KubeContext,
		"--set", fmt.Sprintf("manager.image=%s", strs[0]),
		"--set", fmt.Sprintf("manager.tag=%s", strs[1]),
	}

	cmd := exec.Command("helm", args...)
	cmd.Stderr = os.Stderr

	Expect(cmd.Run()).Should(Succeed())
}

func HelmPurgeRelease(release, ns string) {
	args := []string{
		"delete", release,
		"--namespace", ns,
		"--kube-context", TestContext.KubeContext,
	}
	cmd := exec.Command("helm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	Expect(cmd.Run()).Should(Succeed())
}
