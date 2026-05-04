@Library('jenkins-library') _

microservicePipeline(
    environment: branchToEnvironment(env.BRANCH_NAME),

    test: {
        runTests(services: ['auth-api', 'data-api'])
    },

    build: {
        buildImage(image: 'REGISTRY_PLACEHOLDER/auth-api', context: 'auth-api')
        buildImage(image: 'REGISTRY_PLACEHOLDER/data-api', context: 'data-api')
    }
)
