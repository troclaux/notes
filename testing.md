
# software testing

## types of tests

- functional testing: verifies operations and actions of an application
- non-functional testing: verifies behaviour of application
  - performance
  - security
  - usability

### functional testing

- unit test: individual units of code (functions, methods)
- smoke testing: minimal set of tests that check basic functionality
  - normally applied after new build is deployed
- integration test: how different units of code work together to ensure they integrate correctly
- regression testing: rerun tests after minor changes to ensure functionality still works correctly
- end-to-end test: checks if code works correctly from start to finish (including all interactions and dependencies)

### non-functional testing

- performance test: checks if the code meets performance requirements
- load testing: apply specific load to system to understand its behaviour under specific conditions
  - tests system in normal conditions
- stress testing: how much stress the system can take before it fails
  - tests system in extreme conditions
- security testing: checks system's ability to protect data and maintain functionality during malicious attacks
- scalability testing: tests system's ability to scale up or down efficiently
