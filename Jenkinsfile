pipeline {
    agent any

    stages {
        stage('Clone') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/master']], extensions: [],
                    userRemoteConfigs: [[url: 'https://github.com/kubesphere-sigs/demo-go-http']]])
            }
        }

        stage('Build') {
            steps {
                // generate the image tag name
                script {
                    env.TAGNAME=sh returnStdout: true, script: 'git rev-parse --short  HEAD'
                }

                container('base') {
                    sh '''
                        echo $TAGNAME
                        docker build . -t docker.io/surenpi/test:$TAGNAME
                    '''
                }
            }
        }

        stage('Push Image') {
            steps {
                container('base') {
                    withCredentials([usernamePassword(credentialsId : 'dockerhub' ,passwordVariable : 'PASS' ,usernameVariable : 'USER' ,)]) {
                        sh '''
                            docker login -u$USER -p$PASS docker.io
                            docker push docker.io/surenpi/test:$TAGNAME
                        '''
                    }
                }
            }
        }
    }
}