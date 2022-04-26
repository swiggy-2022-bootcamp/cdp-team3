#!groovy
pipeline {
    agent any
    tools {
        go 'Go'
    }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        APP_NAME = "shopping"
        REGION = "${REGION}"
        ACCESS_KEY = "${AWS_SECRET_ACCESS_KEY}"
        KEY_ID = "${AWS_ACCESS_KEY_ID}"
        JWT_SECRET_KEY = "${JWT_SECRET_KEY}"
    }
    stages {
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
            }
        }

        stage('Build #1') {
            steps {
                echo 'Compiling and building'
                sh 'cd shipping-service && go build'
            }
        }

        stage('Test #1') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd shipping-service && go vet .'
                    echo 'Running test'
                    sh 'cd shipping-service && go test ./...'
                }
            }
        }

        stage('Create .env file #1') {
            steps {
                sh 'cd shipping-service && touch .env'
                sh 'cd shipping-service && echo "REGION=${REGION}" >> .env'
                sh 'cd shipping-service && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd shipping-service && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd shipping-service && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #1') {
             steps {
                 echo 'Building the latest docker image'
                 sh 'cd shipping-service && chmod +x startup.sh && ./startup.sh'
             }
        }
        stage('Run #1.1') {
             steps {
                 echo 'Pushing image to dockerhub'
                 sh 'cd shipping-service && chmod +x push.sh && ./push.sh'
             }
        }
        stage('Build #2') {
            steps {
                echo 'Compiling and building'
                sh 'cd categories-service && go build'
            }
        }

        stage('Test #2') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd categories-service && go vet .'
                    echo 'Running test'
                    sh 'cd categories-service && go test ./...'
                }
            }
        }

        stage('Create .env file #2') {
            steps {
                sh 'cd categories-service && touch .env'
                sh 'cd categories-service && echo "REGION=${REGION}" >> .env'
                sh 'cd categories-service && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd categories-service && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd categories-service && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #2') {
             steps {
                 echo 'Building the latest docker image'
                 sh 'cd categories-service && chmod +x startup.sh && ./startup.sh'
             }
        }
        stage('Run #2.1') {
             steps {
                 echo 'Pushing image to dockerhub'
                 sh 'cd categories-service && chmod +x push.sh && ./push.sh'
             }
        }
        
         stage('Run #') {
             steps {
                 echo 'Running docker compose'
                 sh 'docker compose up -d'
             }
        }


    }
}