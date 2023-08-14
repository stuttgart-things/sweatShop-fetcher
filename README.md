# stuttgart-things/sweatShop-fetcher

sweatShop-fetcher is a gRPC server, which hosts a repository service.

This service takes in a git repository url, credentials and delivers a list of files, fitting the given pattern.
During the initialization of the repository service, existing repositories in the system tmp dir are scanned and save into a repo list.

When a request to list files of a certain pattern in a repository comes in, the repository service first clone the repository into the system tmp dir, if the repository is not cloned yet. Then the service gathers files fitting the pattern and deliver a file list as output.

## LICENSE

<details><summary><b>APACHE 2.0</b></summary>

Copyright 2023 xiaomin.lai.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

</details>

Author Information
------------------
Xiaomin Lai, stuttgart-things 01/2023

Moved to github,
Patrick Hermann, stuttgart-things 08/2023
