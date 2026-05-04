@Library('jenkins-library') _

pipeline {
    agent {
        kubernetes {
            inheritFrom 'skaffold'
        }
    }
    stages {
        stage('Checkout') {
            steps {
                checkoutApp()
            }
        }
        stage('Test') {
            steps {
                runTests(services: ['auth-api', 'data-api'])
            }
        }
        stage('Build') {
            steps {
                buildImage(image: 'REGISTRY_PLACEHOLDER/auth-api', context: 'auth-api')
                buildImage(image: 'REGISTRY_PLACEHOLDER/data-api', context: 'data-api')
            }
        }
        stage('Release') {
            steps {
                releaseApp(environment: branchToEnvironment(env.BRANCH_NAME))
            }
        }
    }
}
