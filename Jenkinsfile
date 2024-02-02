pipeline {
    agent {
        kubernetes {
            yaml """
                apiVersion: v1
                kind: Pod
                spec:
                  dnsConfig:
                    nameservers:
                      - "8.8.8.8"
                  serviceAccountName: jenkins
                  
                  containers:
                  - name: docker
                    image: docker
                    volumeMounts:
                    - name: docker-socket
                      mountPath: /var/run/docker.sock
                    command:
                    - cat
                    tty: true
                    securityContext:
                      privileged: true
                  - name: kubectl
                    image: bitnami/kubectl:latest
                    command:
                    - cat
                    tty: true
                    securityContext:
                      runAsUser: 1000
                  imagePullSecrets:
                    - name: docker
                  volumes:
                  - name: docker-socket
                    hostPath:
                      path: /var/run/docker.sock
            """
        }
    }

    stages {
        stage("Docker login") {
            when {
                expression {
                    return (env.BRANCH_NAME == 'main')
                }
            }
            steps {
                container('docker') {
                    withCredentials([usernamePassword(credentialsId: 'docker-hub', passwordVariable: 'DOCKER_PASSWORD', usernameVariable: 'DOCKER_USERNAME')]) {
                        sh 'docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD'
                    }
                }
            }
        }
        stage("Building and push image to docker hub") {
            when {
                expression {
                    return (env.BRANCH_NAME == 'main')
                }
            }
            steps {
                container('docker') {
                    sh 'docker build -t nckh:latest .'
                    sh 'docker tag nckh toan3082004/nckh:latest'
                    sh 'docker push toan3082004/nckh:latest'
                }
            }
        }
        stage("Deploy to kubernetes") {
            when {
                expression {
                    return (env.BRANCH_NAME == 'main')
                }
            }
            steps {
                container('kubectl') {
                    sh 'kubectl delete deployment/nckh-backend -n backend || true'
                    sh 'kubectl apply -f deploy/deployment.yaml'
                }
            }
        }
    }
}
