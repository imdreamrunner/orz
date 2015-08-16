module.exports = (grunt) ->
    # Load tasks
    grunt.loadNpmTasks 'grunt-contrib-clean'
    grunt.loadNpmTasks 'grunt-contrib-copy'
    grunt.loadNpmTasks 'grunt-contrib-coffee'
    grunt.loadNpmTasks 'grunt-contrib-less'
    grunt.loadNpmTasks 'grunt-bower-task'
    grunt.loadNpmTasks 'grunt-contrib-watch'
    # Init config
    grunt.initConfig
        clean:
            static: ['_static']
            libraries: ['_libraries']
            bower: ['bower_components']
        copy:
            public:
                expand: true
                cwd: 'public'
                src: ['**']
                dest: '_static'
                filter: 'isFile'
            libraries:
                expand: true
                cwd: '_libraries'
                src: ['**']
                dest: '_static/lib'
                filter: 'isFile'
        coffee:
            assets:
                expand: true
                cwd: 'assets'
                src: ['**/*.coffee']
                dest: '_static'
                ext: '.js'
        less:
            assets:
                expand: true
                cwd: 'assets'
                src: ['**/*.less']
                dest: '_static'
                ext: '.css'
        bower:
            install:
                options:
                    targetDir: '_libraries'
                    layout: 'byComponent'
        watch:
            assets:
                files: ['assets/**/*.*']
                tasks:  ['coffee', 'less']
     grunt.registerTask 'default', ['clean:static', 'copy', 'coffee', 'less']
     grunt.registerTask 'install', ['clean:libraries', 'bower']
