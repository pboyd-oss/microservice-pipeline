@Library('jenkins-library') _

microservicePipeline(
    environment: branchToEnvironment(env.BRANCH_NAME),

    test: {
        runTests(services: ['auth-api', 'data-api'])
    },

    build: {
        buildApp()
    }
)
