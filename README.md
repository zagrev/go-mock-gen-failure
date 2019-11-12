# Go mockgen generate mock failure

This project demonstrates the failure of mockgen when re-generating the mocks for a changed interface.
When the unit tests were written and the mocks originally gneerated, everything works as expected.
But once the mocks exists, and the interface changes, then mockgen fails to generate new mocks because the tests fail to compile.


You can check this in this project using the tags "pre-change" and "post-change"
