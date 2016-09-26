Wercker Java Version Step
=============================================

[![wercker status](https://app.wercker.com/status/4d7e381ef860453ca5a3fcb9a347eaeb/s/master "wercker status")](https://app.wercker.com/project/bykey/4d7e381ef860453ca5a3fcb9a347eaeb)

This repository holds the code to create the wercker step from alianza for 
extracting your maven details from the pom and using that to generate the
build information that is stored on disk.  This can be sourced so that 
future steps and pipelines can maintain the same state.  

## Usage

Since this is a step you will want to add the step to your wercker config, you
can find the step [here](https://app.wercker.com/applications/57dc0dc0ae73a701004e9e5b/tab/details/).    

A sample of the step using in wercker is shown below.   

        steps:
          - alianza/versioning

Using the above defaults will read the pom details from pom.xml, and output the
bash source items to `build_source`     

## Local Test

You can run this on your local box against a test by running `./local-test` in 
the working directory of the repo.  

**Note:** This requires that at least go version 1.6 is installed    
