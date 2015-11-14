var grunt = require('grunt');
grunt.loadNpmTasks('grunt-aws-lambda');

grunt.initConfig({
   lambda_invoke: {
      default: {
         options: {
            file_name: 'doorman.js'
         }
      }
   },
   lambda_deploy: {
      default: {
         package: 'doorman',
         arn: 'arn:aws:lambda:us-east-1:152888288016:function:doorman'
      }
   },
   lambda_package: {
      default: {}
   }
});

grunt.registerTask('deploy', ['lambda_package', 'lambda_deploy'])
