#!/usr/bin/env groovy
// The above line is used to trigger correct syntax highlighting.

    // If anything fails, the whole Pipeline stops.
node('docker') {
     echo "Running ${env.BUILD_ID}"
     checkout scm
     stage('Test Build') {
         docker.image('golang:1.16-alpine').inside {
              sh 'mkdir -p /app'
              sh 'cd /app'
              sh 'cp -r ${WORKSPACE}/*.go /app'
              sh 'go build -o ./webapp'
         }
      }

      stage('Test') {
          docker.image('golang:1.16-alpine').inside {
               sh 'apk add build-base'
               sh 'mkdir -p /app'
               sh 'cd /app'
               sh 'cp -r ${WORKSPACE}/go.mod /app'
               sh 'cp -r ${WORKSPACE}/*.go /app'
               sh 'go clean -cache'
               sh 'go test -v -timeout 60s'
          }
      }      
/*      
      stage('Push image') {
          docker.withRegistry('https://008866760928.dkr.ecr.us-east-1.amazonaws.com/tking-capstone', 'ecr:us-east-1:aws-ecr-repo') {
                def myImage = docker.build('webapp:${env.BUILD_ID}')
                myImage.push()
          }
      }

      stage('Run staging container') {
           sshagent(credentials: ['dd0c4ba0-6705-4613-a807-7a3c53296719']) {
                sh 'docker run -p 8080:8080 --name webapp-test --detach 008866760928.dkr.ecr.us-east-1.amazonaws.com/tking-capstone/webapp:${env.BUILD_ID}'
           }
      }
*/
}
