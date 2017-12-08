package testutils

import (
	"net/http"
	"net/http/httptest"
)

func MockServer(b []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		w.Write(b)
		w.Header().Set("Content-Type", "application/json")
	}))
}

func TenantData() []byte {
	return []byte(`
		{
			"data": [
					{
							"attributes": {
									"created-at": "2017-10-11T18:47:27.69333Z",
									"email": "vpavlin@redhat.com",
									"namespaces": [
											{
													"cluster-url": "https://api.free-int.openshift.com",
													"created-at": "2017-10-11T18:47:28.491233Z",
													"name": "vpavlin-jenkins",
													"state": "created",
													"type": "jenkins",
													"updated-at": "2017-10-11T18:47:28.491233Z",
													"version": "2.0.6"
											},
											{
													"cluster-url": "https://api.free-int.openshift.com",
													"created-at": "2017-10-11T18:47:28.513893Z",
													"name": "vpavlin-che",
													"state": "created",
													"type": "che",
													"updated-at": "2017-10-11T18:47:28.513893Z",
													"version": "2.0.6"
											},
											{
													"cluster-url": "https://api.free-int.openshift.com",
													"created-at": "2017-10-11T18:47:28.56173Z",
													"name": "vpavlin-run",
													"state": "created",
													"type": "run",
													"updated-at": "2017-10-11T18:47:28.56173Z",
													"version": "2.0.6"
											},
											{
													"cluster-url": "https://api.free-int.openshift.com",
													"created-at": "2017-10-11T18:47:28.604475Z",
													"name": "vpavlin",
													"state": "created",
													"type": "user",
													"updated-at": "2017-10-11T18:47:28.604475Z",
													"version": "1.0.91"
											},
											{
													"cluster-url": "https://api.free-int.openshift.com",
													"created-at": "2017-10-11T18:47:28.763171Z",
													"name": "vpavlin-stage",
													"state": "created",
													"type": "stage",
													"updated-at": "2017-10-11T18:47:28.763171Z",
													"version": "2.0.6"
											}
									],
									"profile": "free"
							},
							"id": "2e15e957-0366-4802-bf1e-0d6fe3f11bb6",
							"type": "tenants"
					}
			],
			"meta": {
					"totalCount": 1
			}
	}`)
}

func TenantData2() []byte {
	return []byte(`
		{
			"data": [
					{
							"attributes": {
									"created-at": "2017-10-11T18:47:27.69333Z",
									"email": "foo@redhat.com",
									"namespaces": [
											{
													"cluster-url": "https://api.free-int.openshift.com",
													"created-at": "2017-10-11T18:47:28.491233Z",
													"name": "foo-jenkins",
													"state": "created",
													"type": "jenkins",
													"updated-at": "2017-10-11T18:47:28.491233Z",
													"version": "2.0.6"
											},
											{
													"cluster-url": "https://api.free-int.openshift.com",
													"created-at": "2017-10-11T18:47:28.513893Z",
													"name": "foo-che",
													"state": "created",
													"type": "che",
													"updated-at": "2017-10-11T18:47:28.513893Z",
													"version": "2.0.6"
											},
											{
													"cluster-url": "https://api.free-int.openshift.com",
													"created-at": "2017-10-11T18:47:28.56173Z",
													"name": "foo-run",
													"state": "created",
													"type": "run",
													"updated-at": "2017-10-11T18:47:28.56173Z",
													"version": "2.0.6"
											},
											{
													"cluster-url": "https://api.free-int.openshift.com",
													"created-at": "2017-10-11T18:47:28.604475Z",
													"name": "foo",
													"state": "created",
													"type": "user",
													"updated-at": "2017-10-11T18:47:28.604475Z",
													"version": "1.0.91"
											},
											{
													"cluster-url": "https://api.free-int.openshift.com",
													"created-at": "2017-10-11T18:47:28.763171Z",
													"name": "foo-stage",
													"state": "created",
													"type": "stage",
													"updated-at": "2017-10-11T18:47:28.763171Z",
													"version": "2.0.6"
											}
									],
									"profile": "free"
							},
							"id": "b53ba1e5-66dd-4dcf-af83-9c28aed601ad",
							"type": "tenants"
					}
			],
			"meta": {
					"totalCount": 1
			}
	}`)
}

func DCJenkinsDataIdle() []byte {
	return []byte(`
		{
			"apiVersion": "v1",
			"kind": "DeploymentConfig",
			"metadata": {
					"annotations": {
							"configmap.fabric8.io/update-on-change": "jenkins",
							"fabric8.io/git-branch": "release-v4.0.41",
							"fabric8.io/git-commit": "2b99f551ea8d500df90ea4e22e3348276ac6cae8",
							"fabric8.io/iconUrl": "https://cdn.rawgit.com/fabric8-services/tenant-jenkins/tenant-jenkins-apps/jenkins/master/apps/jenkins/src/main/fabric8/icon.svg",
							"fabric8.io/metrics-path": "dashboard/file/kubernetes-pods.json/?var-project=jenkins\u0026var-version=4.0.41",
							"fabric8.io/scm-con-url": "scm:git:git@github.com:fabric8-services/tenant-jenkins.git/tenant-jenkins-apps/jenkins",
							"fabric8.io/scm-devcon-url": "scm:git:git@github.com:fabric8-services/tenant-jenkins.git/tenant-jenkins-apps/jenkins",
							"fabric8.io/scm-tag": "fabric8-services/tenant-jenkins-1.0.0",
							"fabric8.io/scm-url": "http://github.com/fabric8-services/tenant-jenkins/tenant-jenkins-apps/jenkins",
							"fabric8.io/target-platform": "openshift",
							"idling.alpha.openshift.io/idled-at": "2017-12-04T18:06:54.273889446Z",
							"idling.alpha.openshift.io/previous-scale": "1",
							"kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"v1\",\"kind\":\"DeploymentConfig\",\"metadata\":{\"annotations\":{\"configmap.fabric8.io/update-on-change\":\"jenkins\",\"fabric8.io/git-branch\":\"release-v4.0.41\",\"fabric8.io/git-commit\":\"2b99f551ea8d500df90ea4e22e3348276ac6cae8\",\"fabric8.io/iconUrl\":\"https://cdn.rawgit.com/fabric8-services/tenant-jenkins/tenant-jenkins-apps/jenkins/master/apps/jenkins/src/main/fabric8/icon.svg\",\"fabric8.io/metrics-path\":\"dashboard/file/kubernetes-pods.json/?var-project=jenkins\\u0026var-version=4.0.41\",\"fabric8.io/scm-con-url\":\"scm:git:git@github.com:fabric8-services/tenant-jenkins.git/tenant-jenkins-apps/jenkins\",\"fabric8.io/scm-devcon-url\":\"scm:git:git@github.com:fabric8-services/tenant-jenkins.git/tenant-jenkins-apps/jenkins\",\"fabric8.io/scm-tag\":\"fabric8-services/tenant-jenkins-1.0.0\",\"fabric8.io/scm-url\":\"http://github.com/fabric8-services/tenant-jenkins/tenant-jenkins-apps/jenkins\",\"fabric8.io/target-platform\":\"openshift\",\"maven.fabric8.io/source-url\":\"jar:file:/home/jenkins/workspace/bric8-tenant-jenkins_master-AGRITQKDOGHX6ZQLEFCNJVZP7VWPA5P7B6D2RTTKZ2KN4XI2LSCA/apps/jenkins/target/jenkins-4.0.41.jar!/META-INF/fabric8/openshift.yml\"},\"labels\":{\"app\":\"jenkins\",\"group\":\"io.fabric8.tenant.apps\",\"provider\":\"fabric8\",\"version\":\"4.0.41\"},\"name\":\"jenkins\",\"namespace\":\"vpavlin-jenkins\"},\"spec\":{\"replicas\":1,\"revisionHistoryLimit\":2,\"selector\":{\"app\":\"jenkins\",\"group\":\"io.fabric8.tenant.apps\",\"provider\":\"fabric8\"},\"strategy\":{\"recreateParams\":{\"timeoutSeconds\":10000},\"type\":\"Recreate\"},\"template\":{\"metadata\":{\"annotations\":{\"configmap.fabric8.io/update-on-change\":\"jenkins\",\"fabric8.io/git-branch\":\"release-v4.0.41\",\"fabric8.io/git-commit\":\"2b99f551ea8d500df90ea4e22e3348276ac6cae8\",\"fabric8.io/iconUrl\":\"https://cdn.rawgit.com/fabric8-services/tenant-jenkins/tenant-jenkins-apps/jenkins/master/apps/jenkins/src/main/fabric8/icon.svg\",\"fabric8.io/metrics-path\":\"dashboard/file/kubernetes-pods.json/?var-project=jenkins\\u0026var-version=4.0.41\",\"fabric8.io/scm-con-url\":\"scm:git:git@github.com:fabric8-services/tenant-jenkins.git/tenant-jenkins-apps/jenkins\",\"fabric8.io/scm-devcon-url\":\"scm:git:git@github.com:fabric8-services/tenant-jenkins.git/tenant-jenkins-apps/jenkins\",\"fabric8.io/scm-tag\":\"fabric8-services/tenant-jenkins-1.0.0\",\"fabric8.io/scm-url\":\"http://github.com/fabric8-services/tenant-jenkins/tenant-jenkins-apps/jenkins\",\"fabric8.io/target-platform\":\"openshift\"},\"labels\":{\"app\":\"jenkins\",\"group\":\"io.fabric8.tenant.apps\",\"provider\":\"fabric8\",\"version\":\"4.0.41\"}},\"spec\":{\"containers\":[{\"env\":[{\"name\":\"PROJECT_NAMESPACE\",\"value\":\"vpavlin\"},{\"name\":\"KUBERNETES_NAMESPACE\",\"valueFrom\":{\"fieldRef\":{\"fieldPath\":\"metadata.namespace\"}}},{\"name\":\"GIT_COMMITTER_EMAIL\",\"value\":\"fabric8@googlegroups.com\"},{\"name\":\"GIT_COMMITTER_NAME\",\"value\":\"fabric8\"},{\"name\":\"OPENSHIFT_ENABLE_OAUTH\",\"value\":\"true\"},{\"name\":\"OPENSHIFT_ENABLE_REDIRECT_PROMPT\",\"value\":\"true\"},{\"name\":\"KUBERNETES_TRUST_CERTIFICATES\",\"value\":\"true\"},{\"name\":\"KUBERNETES_MASTER\",\"value\":\"https://kubernetes.default:443\"},{\"name\":\"JAVA_GC_OPTS\",\"value\":\"-XX:+UseParallelGC -XX:MinHeapFreeRatio=5 -XX:MaxHeapFreeRatio=10 -XX:GCTimeRatio=4 -XX:AdaptiveSizePolicyWeight=90\"},{\"name\":\"JAVA_OPTS\",\"value\":\"-XX:+UnlockExperimentalVMOptions -XX:+UseCGroupMemoryLimitForHeap -Dsun.zip.disableMemoryMapping=true\"},{\"name\":\"OPENSHIFT_JENKINS_JVM_ARCH\",\"value\":\"i686\"},{\"name\":\"CONTAINER_INITIAL_PERCENT\",\"value\":\"0.07\"},{\"name\":\"CONTAINTER_INITIAL_PERCENT\",\"value\":\"something-non-empty\"},{\"name\":\"JENKINS_OPTS\",\"value\":\"-Dgroovy.use.classvalue=true\"},{\"name\":\"JAVA_HOME\",\"value\":\"/etc/alternatives/java_sdk\"},{\"name\":\"MAVEN_OPTS\",\"value\":\"-Dorg.slf4j.simpleLogger.log.org.apache.maven.cli.transfer.Slf4jMavenTransferListener=warn\"},{\"name\":\"RECOMMENDER_API_TOKEN\",\"valueFrom\":{\"secretKeyRef\":{\"key\":\"token\",\"name\":\"jenkins-recommender-api-token\"}}}],\"image\":\"fabric8/jenkins-openshift:vb6fc097\",\"imagePullPolicy\":\"IfNotPresent\",\"livenessProbe\":{\"failureThreshold\":30,\"httpGet\":{\"path\":\"/login\",\"port\":8080},\"initialDelaySeconds\":420,\"timeoutSeconds\":10},\"name\":\"jenkins\",\"ports\":[{\"containerPort\":50000,\"name\":\"slave\"},{\"containerPort\":8080,\"name\":\"http\"}],\"readinessProbe\":{\"httpGet\":{\"path\":\"/login\",\"port\":8080},\"initialDelaySeconds\":10,\"timeoutSeconds\":10},\"resources\":{\"limits\":{\"cpu\":\"2\",\"memory\":\"512Mi\"},\"requests\":{\"cpu\":\"0\"}},\"volumeMounts\":[{\"mountPath\":\"/var/lib/jenkins\",\"name\":\"jenkins-home\",\"readOnly\":false},{\"mountPath\":\"/opt/openshift/configuration/\",\"name\":\"jenkins-config\"}]}],\"serviceAccountName\":\"jenkins\",\"volumes\":[{\"name\":\"jenkins-home\",\"persistentVolumeClaim\":{\"claimName\":\"jenkins-home\"}},{\"configMap\":{\"name\":\"jenkins\"},\"name\":\"jenkins-config\"}]}},\"triggers\":[{\"type\":\"ConfigChange\"}]}}\n",
							"maven.fabric8.io/source-url": "jar:file:/home/jenkins/workspace/bric8-tenant-jenkins_master-AGRITQKDOGHX6ZQLEFCNJVZP7VWPA5P7B6D2RTTKZ2KN4XI2LSCA/apps/jenkins/target/jenkins-4.0.41.jar!/META-INF/fabric8/openshift.yml"
					},
					"creationTimestamp": "2017-11-30T16:00:06Z",
					"generation": 19,
					"labels": {
							"app": "jenkins",
							"group": "io.fabric8.tenant.apps",
							"provider": "fabric8",
							"version": "4.0.41"
					},
					"name": "jenkins",
					"namespace": "vpavlin-jenkins",
					"resourceVersion": "136008995",
					"selfLink": "/oapi/v1/namespaces/vpavlin-jenkins/deploymentconfigs/jenkins",
					"uid": "8890f892-d5e7-11e7-9451-0a46c474dfe0"
			},
			"spec": {
					"replicas": 0        
			},
			"status": {
					"availableReplicas": 0,
					"conditions": [
							{
									"lastTransitionTime": "2017-11-30T16:50:19Z",
									"lastUpdateTime": "2017-11-30T16:50:19Z",
									"message": "replication controller \"jenkins-2\" successfully rolled out",
									"reason": "NewReplicationControllerAvailable",
									"status": "True",
									"type": "Progressing"
							},
							{
									"lastTransitionTime": "2017-12-04T18:06:55Z",
									"lastUpdateTime": "2017-12-04T18:06:55Z",
									"message": "Deployment config does not have minimum availability.",
									"status": "False",
									"type": "Available"
							}
					],
					"details": {
							"causes": [
									{
											"type": "Manual"
									}
							],
							"message": "manual change"
					},
					"latestVersion": 2,
					"observedGeneration": 19,
					"replicas": 0,
					"unavailableReplicas": 0,
					"updatedReplicas": 0
			}
	}	
		`)
}