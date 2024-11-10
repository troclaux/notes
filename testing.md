
# software testing

## types of tests

- functional testing: verifies operations and actions of an application
  - defines:
    - software capabilities
    - how the system reacts to specific input
    - how the system behaves in specific scenarios
    - what the system cannot do
- non-functional testing: verifies behaviour of application
  - performance
  - security
  - usability

- white box test: uses knowledge about code to test correctness of software at code level
- black box test: verifies software behaviour without internal code knowledge, focus on input and output

- system testing: checks if the entire system meets functional and non-functional requirements
- validation test: checks if software behaves according to user expectation and requirements

| aspect              | alpha testing                         | beta testing                          |
|---------------------|--------------------------------------|---------------------------------------|
| conducted by    | internal team (developers/testers)  | external users                        |
| stage of testing| early stage (pre-release)           | late stage (pre-launch)              |
| environment     | controlled environment               | real-world environment                |
| goal            | identify bugs and issues             | validate the product and gather user feedback |
| user feedback    | immediate feedback for quick fixes   | broader feedback on usability and experience |

### functional testing

- unit test: smallest units of code (functions, methods)
- smoke testing: preliminary set of tests that check basic functionality
  - normally applied after new build is deployed
- integration test: ensures the software components or functions work together correctly
- regression testing: check if new features break or degrade functionality
- end-to-end test: checks if code works correctly from start to finish (including all interactions and dependencies)

### non-functional testing

- performance test: checks if the code meets performance requirements
- load testing: apply specific load to system to understand its behaviour under specific conditions
  - tests system in normal conditions
- stress testing: how much stress the system can take before it fails
  - tests system in extreme conditions
- security testing: checks system's ability to protect data and maintain functionality during malicious attacks
- scalability testing: tests system's ability to scale up or down efficiently
