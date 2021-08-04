#!/usr/bin/env groovy
// The above line is used to trigger correct syntax highlighting.

pipeline {
    // Lets Jenkins use Docker for us later.
    agent { label 'docker' }

    // If anything fails, the whole Pipeline stops.
    stages {
         stage('Build Test') {
               steps {
                     script {
                           docker.image('golang:1.16-alpine').inside {
                                sh 'mkdir -p /app'
                                sh 'cd /app'
                                sh 'cp -r ${WORKSPACE}/* /app'
                                sh 'go build -o ./webapp'
                           }
                     }
               }
          }
    }
}
/*				
        stage('Build') {   
            // Use golang.
            agent { docker { image 'golang:1.16-alpine' } }

            steps {                                           
                // Create our project directory.
                sh 'mkdir -p /app'
                sh 'cd /app'

                // Copy all files in our Jenkins workspace to our project directory.                
                sh 'cp -r ${WORKSPACE}/* /app'

                // Build the app.
                sh 'go build -o ./webapp'               
            }            
        }

        stage('Test') {
            // Use golang.
            agent { docker { image 'golang:1.16-alpine' } }

            steps {                 
                // Create our project directory.
                sh 'mkdir -p /app'
                sh 'cd /app'

                // Copy all files in our Jenkins workspace to our project directory.                
                sh 'cp -r ${WORKSPACE}/* /app'

		// Remove cached test results.
                sh 'go clean -cache'

                // Run Unit Tests. 
                sh 'go test -v -timeout 60s'
            }
        }      
    }
}
*/
/*        stage('Docker') {         
            environment {
                // Extract the username and password of our credentials into "DOCKER_CREDENTIALS_USR" and "DOCKER_CREDENTIALS_PSW".
                // (NOTE 1: DOCKER_CREDENTIALS will be set to "your_username:your_password".)
                // The new variables will always be YOUR_VARIABLE_NAME + _USR and _PSW.
                // (NOTE 2: You can't print credentials in the pipeline for security reasons.)
                DOCKER_CREDENTIALS = credentials('my-docker-credentials-id')
            }

            steps {                           
                // Use a scripted pipeline.
                script {
                    node ('docker') {
                        def app

                        stage('Clone repository') {
                            checkout scm
                        }

                        stage('Build image') {                            
                            app = docker.build("${env.DOCKER_CREDENTIALS_USR}/my-project-img")
                        }

                        stage('Push image') {  
                            // Use the Credential ID of the Docker Hub Credentials we added to Jenkins.
                            docker.withRegistry('https://registry.hub.docker.com', 'my-docker-credentials-id') {                                
                                // Push image and tag it with our build number for versioning purposes.
                                app.push("${env.BUILD_NUMBER}")                      

                                // Push the same image and tag it as the latest version (appears at the top of our version list).
                                app.push("latest")
                            }
                        }              
                    }
                }
            }
        }
    }

    post {
        always {
            // Clean up our workspace.
            deleteDir()
        }
    }
}  */ 
