/*******************************************************************************
 * Copyright 2024 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package contracts

type ContentType string

const (
	ContentTypeJSON ContentType = "application/json"
)

type NetType string

const (
	Mainnet    NetType = "mainnet"
	Testnet    NetType = "testnet"
	Previewnet NetType = "previewnet"
	Local      NetType = "local"
)

func (t NetType) Validate() bool {
	if t == Mainnet || t == Testnet || t == Previewnet || t == Local {
		return true
	}
	return false
}

type HashType string

const (
	MD5Hash    HashType = "md5"
	SHA256Hash HashType = "sha256"
	NoHash     HashType = "none"
)

func (t HashType) Validate() bool {
	if t == MD5Hash || t == SHA256Hash || t == NoHash {
		return true
	}
	return false
}

type KeyAlgorithm string

const (
	KeyEd25519        KeyAlgorithm = "ed25519"
	KeyEcdsaX509      KeyAlgorithm = "ecdsa-x509"
	KeyEcdsaSecp256k1 KeyAlgorithm = "ecdsa-secp256k1"
)

func (k KeyAlgorithm) Validate() bool {
	if k == KeyEd25519 || k == KeyEcdsaX509 || k == KeyEcdsaSecp256k1 {
		return true
	}
	return false
}

type StreamType string

const (
	ConsoleStream StreamType = "console"
	MockStream    StreamType = "mock"
	MqttStream    StreamType = "mqtt"
	PravegaStream StreamType = "pravega" // Currently unsupported but indicating extension point
	HederaStream  StreamType = "hedera"
)

func (t StreamType) Validate() bool {
	if t == MockStream || t == MqttStream || t == PravegaStream || t == ConsoleStream || t == HederaStream {
		return true
	}
	return false
}

type AnnotationType string

const (
	AnnotationPKI     AnnotationType = "pki"
	AnnotationPKIHttp AnnotationType = "pki-http"
	AnnotationSource  AnnotationType = "src"
	AnnotationTLS     AnnotationType = "tls"
	AnnotationTPM     AnnotationType = "tpm"
	// The AnnotationSourceCode, AnnotationChecksum, and AnnotationVulnerability values are used by the scoring apps, they are for CI/CD annotators defined in alvarium-sdk-java project.
	AnnotationSourceCode    AnnotationType = "source-code"
	AnnotationChecksum      AnnotationType = "checksum"
	AnnotationVulnerability AnnotationType = "vulnerability"
	AnnotationSBOM          AnnotationType = "sbom"
)

func (t AnnotationType) Validate() bool {
	switch t {
	case AnnotationPKI, AnnotationTLS, AnnotationTPM, AnnotationSource, AnnotationPKIHttp, AnnotationSourceCode, AnnotationChecksum, AnnotationVulnerability:
		return true
	default:
		return false
	}
}

type DerivedComponent string

const (
	Method      DerivedComponent = "@method"
	TargetURI   DerivedComponent = "@target-uri"
	Authority   DerivedComponent = "@authority"
	Scheme      DerivedComponent = "@scheme"
	Path        DerivedComponent = "@path"
	Query       DerivedComponent = "@query"
	QueryParams DerivedComponent = "@query-params"
)

const (
	// HttpRequestKey is the key used to reference the value within the incoming Context that corresponds to the request we need to validate.
	HttpRequestKey  string = "HttpRequestKey"
	ContentLength   string = "Content-Length"
	HttpContentType string = "Content-Type"
)

func (d DerivedComponent) Validate() bool {
	if d == Method || d == Authority || d == TargetURI || d == Scheme || d == Path || d == Query || d == QueryParams {
		return true
	}
	return false
}

type LayerType string

const (
	Application LayerType = "app"
	CiCd        LayerType = "cicd"
	Os          LayerType = "os"
	Host        LayerType = "host"
)

func (l LayerType) Validate() bool {
	switch l {
	case Application, CiCd, Os, Host:
		return true
	default:
		return false
	}
}
