
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

- white box testing: uses knowledge about code to test correctness of software at code level
- black box testing: verifies software behaviour without internal code knowledge, focus on input and output

- system testing: checks if the entire system meets functional and non-functional requirements
- validation test: checks if the software behaves according to user expectations and requirements

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
- end-to-end test: checks if the code works correctly from start to finish (including all interactions and dependencies)

### non-functional testing

- performance test: checks if the code meets performance requirements
- load testing: applies a specific load to the system to understand its behavior under defined conditions
  - tests system in normal conditions
- stress testing: how much stress the system can take before it fails
  - tests system in extreme conditions
- security testing: checks system's ability to protect data and maintain functionality during malicious attacks
- scalability testing: tests system's ability to scale up or down efficiently

## A/B testing

> experiments used to compare two versions of something to see which performs better

- also called split tests

1. randomly divide users into two groups

- group A sees version A (the control, which is the current version)
- group B sees version B (the variant, which is the new or modified version)

2. then measure a key metric (e.g. click rate, purchase rate, time on page) for both groups
