@Library('jenkins-library') _

pipeline {
    agent {
        kubernetes {
            cloud env.TUXGRID_BUILD_CLOUD
            inheritFrom 'base'
        }
    }

    stages {
        stage('Checkout') {
            steps { checkout scm }
        }

        stage('Test') {
            steps {
                script {
                    runTests(services: ['auth-api', 'data-api'])
                }
            }
        }

        stage('Build') {
            steps {
                script {
                    buildApp()
                }
            }
        }
    }

    post {
        always {
            junit allowEmptyResults: true, testResults: '**/test-results/*.xml'
        }
    }
}
