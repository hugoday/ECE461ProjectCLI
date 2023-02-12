# ECE 461 Project CLI

Project Members:
Hugo Day, Luke Diehm, B.J. Pemberton, Aubrey Gatewood

Project Description:
ECE 461 - CLI implementation for ranking of Github and NPMJS repositories/packages.

Features:
Our CLI takes in a file with a list of URLs.
We parse those URLs and perform metric calculations to score the repositories.
Metrics include: Correctness, Bus Factor, Responsiveness, Ramp-Up Time, and License Compatibility.

Requirements:
Ramp-Up Time is dependent on cloc - https://github.com/AlDanial/cloc
We are required to accept NPMJS Package URLs and Github Repository URLs - Currently only accept Github URLs.
We are required to take the URL and output a net score based on the five calculated metrics.
We are required to use at least 30% of a specific language from a list of languages - we chose Go.
We are required to have a test suite that gets 80% code coverage - We have a test suit, need to reach 80% code coverage.
We are required to use Github GraphQl and Rest APIs to calculate one metric each - we have done this (GraphQl API - Responsiveness, REST API - Correctness).
We are required to clone the repository to calculate at least one metric - we have done this (Bus Factor, License Compatibility, Ramp-Up Time).
We are required to check if the repository is compatbile with LGPLv2.1 license - we have done this (License Compatibility - checks if this is the license the repository uses).
We are required to send our final output ranked list to stdout - we have done this in NDJSON format.
We are using a GitHub Token in environment variables to run REST and GraphQL APIs.

