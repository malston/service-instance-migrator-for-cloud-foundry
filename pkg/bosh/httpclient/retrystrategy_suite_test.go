/* 
 *  Copyright 2022 VMware, Inc.
 *  
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *  http://www.apache.org/licenses/LICENSE-2.0
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package httpclient_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRetrystrategy(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Retrystrategy Suite")
}

type simpleRetryable struct {
	attemptOutputs []attemptOutput
	Attempts       int
}

type attemptOutput struct {
	ShouldRetry bool
	AttemptErr  error
}

func newSimpleRetryable(attemptOutputs []attemptOutput) *simpleRetryable {
	return &simpleRetryable{
		attemptOutputs: attemptOutputs,
	}
}

func (r *simpleRetryable) Attempt() (bool, error) {
	r.Attempts++

	if len(r.attemptOutputs) > 0 {
		attemptOutput := r.attemptOutputs[0]
		r.attemptOutputs = r.attemptOutputs[1:]
		return attemptOutput.ShouldRetry, attemptOutput.AttemptErr
	}

	return true, nil
}