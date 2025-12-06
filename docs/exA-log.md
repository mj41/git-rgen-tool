# Example rgen run

Source repository gl-exA-src [commit 880d35d /assets tree](https://github.com/mj41/gl-exA-src/tree/880d35db40bf80dc967596fda31872deea665c76/assets) (on GitHub)

## Script output
```
Source repository gl-exA-src latest commit:
880d35d Change start date as we have too many commits already

Repository generation:
Removing existing gl-exA directory and regenerating with git-rgen...
Commands:
  rm -rf ../gl-exA
  go run ./cmd/rgen --conf ../gl-exA-src/rgen-conf.json

Initialized empty Git repository in /home/mj/work-stai/gl-exA/.git/
Processing 001-init-repo...
[main (root-commit) 39ef7fc] Repo initialization and simple readme
 1 file changed, 3 insertions(+)
 create mode 100644 readme.md
Processing 002-commit-spec...
[main 8f71e3e] Adding simple spec my-tiny-spec.md
 2 files changed, 16 insertions(+)
 create mode 100644 docs/my-tiny-spec.md
Processing 003-impl-v001...
[main c93c24e] Implement v0.0.1
 1 file changed, 3 insertions(+)
 create mode 100644 pkg/tinylang/tinylang.go
Processing 004-more-docs...
[main a3d7f77] Change docs
 1 file changed, 6 insertions(+), 2 deletions(-)
Processing 005-update-spec...
[main 54949bc] Update spec with intro
 1 file changed, 2 insertions(+)
Processing 006-add-test-target...
[main 654caff] Add test target
 1 file changed, 5 insertions(+)
 create mode 100644 gl-test-data/test_target.txt
Processing 007-add-test-link...
[main 88546a7] Add test link
 1 file changed, 1 insertion(+)
 create mode 100644 gl-test-data/test_link.md
Processing 008-insert-before-target...
[main 6fc4ac2] Insert 2 lines before target line
 2 files changed, 8 insertions(+)
 create mode 100644 gl-test.json
Processing 009-insert-after-target...
[main 3756ccf] Insert 3 lines after target line
 2 files changed, 4 insertions(+), 1 deletion(-)
Processing 010-delete-before-target...
[main db4a3ae] Delete 1 line before target line
 2 files changed, 2 insertions(+), 3 deletions(-)
Processing 011-minor-modification...
[main b8feebd] Minor modification to target line (~20%)
 2 files changed, 2 insertions(+), 2 deletions(-)
Processing 012-major-modification...
[main e545fcd] Major modification to target line (~60%)
 2 files changed, 3 insertions(+), 3 deletions(-)
Processing 013-rename-target...
[main f77a101] Rename target file
 2 files changed, 1 insertion(+), 1 deletion(-)
 rename gl-test-data/{test_target.txt => renamed_target.txt} (100%)
Processing 014-move-to-subdir...
[main 08ecf09] Move target file to subdirectory
 2 files changed, 1 insertion(+), 1 deletion(-)
 rename gl-test-data/{ => subdir}/renamed_target.txt (100%)
Processing 015-multiple-links...
[main 1244c7b] Add second link file pointing to same target
 3 files changed, 12 insertions(+), 2 deletions(-)
 create mode 100644 gl-test-data/test_link2.md
Processing 016-relative-parent-link...
[main 6a02c46] Add nested relative path link using ../
 2 files changed, 7 insertions(+), 1 deletion(-)
 create mode 100644 gl-test-data/nested/deep_link.md
Processing 017-manual-link-update...
[main 6844f71] Manual link update to reset origin
 2 files changed, 3 insertions(+), 3 deletions(-)
Processing 018-heuristic-recovery...
[main af9d679] Delete target line and add similar line nearby
 2 files changed, 2 insertions(+), 2 deletions(-)
Processing 019-broken-link...
[main 9ab56ea] Delete target line without replacement (broken link)
 3 files changed, 4 insertions(+), 5 deletions(-)
Processing 020-multi-target-setup...
[main 0dd3fd3] Add multi-commit tracking target file
 2 files changed, 23 insertions(+)
 create mode 100644 gl-test-data/multi_link.md
 create mode 100644 gl-test-data/multi_target.txt
Processing 021-multi-insert...
[main f0c12d8] Insert 5 lines at L5 in multi_target
 1 file changed, 5 insertions(+)
Processing 022-multi-delete...
[main c2d8c6a] Delete 2 lines at beginning of multi_target
 1 file changed, 2 deletions(-)
Processing 023-multi-modify...
[main 749c6d2] Minor modification to multi_target line
 1 file changed, 1 insertion(+), 1 deletion(-)
Processing 024-absolute-path-link...
[main 820a162] Add absolute path link
 1 file changed, 5 insertions(+)
 create mode 100644 gl-test-data/absolute_link.md
Processing 025-link-in-go-file...
[main 730388d] Add link in Go source file comment
 1 file changed, 16 insertions(+)
 create mode 100644 gl-test-data/code_with_link.go
Processing 026-whitespace-changes...
[main c288b86] Test whitespace changes around target
 1 file changed, 4 insertions(+), 1 deletion(-)
Processing 027-large-file-setup...
[main 07d7c83] Add large target file for performance testing
 2 files changed, 109 insertions(+)
 create mode 100644 gl-test-data/large_links.md
 create mode 100644 gl-test-data/large_target.txt
Processing 028-large-file-changes...
[main 88752c0] Modify large file with multiple insertions and deletions
 1 file changed, 8 insertions(+), 3 deletions(-)
Processing 029-modified-source-file...
[main 5a8abcd] Add link in file that will be modified in working directory
 2 files changed, 6 insertions(+)
 create mode 100644 gl-test-data/modify_source_link.md
 create mode 100644 gl-test-data/modify_source_target.txt
Processing 030-shift-for-modified-test...
[main b325807] Shift target line for modified source file test
 2 files changed, 29 insertions(+), 7 deletions(-)

Generated repository full log:

 39ef7fc Repo initialization and simple readme
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Sun Jun 22 08:16:32 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Sun Jun 22 09:18:36 2025 +0000

 readme.md | 3 +++
 1 file changed, 3 insertions(+)

 8f71e3e Adding simple spec my-tiny-spec.md
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Mon Jun 23 09:17:33 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Mon Jun 23 10:19:37 2025 +0000

 docs/my-tiny-spec.md | 8 ++++++++
 readme.md            | 8 ++++++++
 2 files changed, 16 insertions(+)

 c93c24e Implement v0.0.1
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Tue Jun 24 10:18:34 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Tue Jun 24 11:20:38 2025 +0000

 pkg/tinylang/tinylang.go | 3 +++
 1 file changed, 3 insertions(+)

 a3d7f77 Change docs
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Wed Jun 25 11:19:35 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Wed Jun 25 12:21:39 2025 +0000

 docs/my-tiny-spec.md | 8 ++++++--
 1 file changed, 6 insertions(+), 2 deletions(-)

 54949bc Update spec with intro
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Thu Jun 26 12:20:36 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Thu Jun 26 13:22:40 2025 +0000

 docs/my-tiny-spec.md | 2 ++
 1 file changed, 2 insertions(+)

 654caff Add test target
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Fri Jun 27 13:21:37 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Fri Jun 27 14:23:41 2025 +0000

 gl-test-data/test_target.txt | 5 +++++
 1 file changed, 5 insertions(+)

 88546a7 Add test link
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Sat Jun 28 14:22:38 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Sat Jun 28 15:24:42 2025 +0000

 gl-test-data/test_link.md | 1 +
 1 file changed, 1 insertion(+)

 6fc4ac2 Insert 2 lines before target line
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Sun Jun 29 15:23:39 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Sun Jun 29 16:25:43 2025 +0000

 gl-test-data/test_target.txt | 2 ++
 gl-test.json                 | 6 ++++++
 2 files changed, 8 insertions(+)

 3756ccf Insert 3 lines after target line
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Mon Jun 30 16:24:40 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Mon Jun 30 17:26:44 2025 +0000

 gl-test-data/test_target.txt | 3 +++
 gl-test.json                 | 2 +-
 2 files changed, 4 insertions(+), 1 deletion(-)

 db4a3ae Delete 1 line before target line
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Tue Jul 1 17:25:41 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Tue Jul 1 18:27:45 2025 +0000

 gl-test-data/test_target.txt | 1 -
 gl-test.json                 | 4 ++--
 2 files changed, 2 insertions(+), 3 deletions(-)

 b8feebd Minor modification to target line (~20%)
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Wed Jul 2 18:26:42 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Wed Jul 2 19:28:46 2025 +0000

 gl-test-data/test_target.txt | 2 +-
 gl-test.json                 | 2 +-
 2 files changed, 2 insertions(+), 2 deletions(-)

 e545fcd Major modification to target line (~60%)
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Thu Jul 3 19:27:43 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Thu Jul 3 20:29:47 2025 +0000

 gl-test-data/test_target.txt | 2 +-
 gl-test.json                 | 4 ++--
 2 files changed, 3 insertions(+), 3 deletions(-)

 f77a101 Rename target file
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Fri Jul 4 20:28:44 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Fri Jul 4 21:30:48 2025 +0000

 gl-test-data/{test_target.txt => renamed_target.txt} | 0
 gl-test.json                                         | 2 +-
 2 files changed, 1 insertion(+), 1 deletion(-)

 08ecf09 Move target file to subdirectory
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Sat Jul 5 21:29:45 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Sat Jul 5 22:31:49 2025 +0000

 gl-test-data/{ => subdir}/renamed_target.txt | 0
 gl-test.json                                 | 2 +-
 2 files changed, 1 insertion(+), 1 deletion(-)

 1244c7b Add second link file pointing to same target
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Sun Jul 6 22:30:46 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Sun Jul 6 23:32:50 2025 +0000

 gl-test-data/test_link.md  | 2 ++
 gl-test-data/test_link2.md | 5 +++++
 gl-test.json               | 7 +++++--
 3 files changed, 12 insertions(+), 2 deletions(-)

 6a02c46 Add nested relative path link using ../
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Mon Jul 7 23:31:47 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Tue Jul 8 00:33:51 2025 +0000

 gl-test-data/nested/deep_link.md | 5 +++++
 gl-test.json                     | 3 ++-
 2 files changed, 7 insertions(+), 1 deletion(-)

 6844f71 Manual link update to reset origin
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Wed Jul 9 00:32:48 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Wed Jul 9 01:34:52 2025 +0000

 gl-test-data/test_link.md | 2 +-
 gl-test.json              | 4 ++--
 2 files changed, 3 insertions(+), 3 deletions(-)

 af9d679 Delete target line and add similar line nearby
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Thu Jul 10 01:33:49 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Thu Jul 10 02:35:53 2025 +0000

 gl-test-data/subdir/renamed_target.txt | 2 +-
 gl-test.json                           | 2 +-
 2 files changed, 2 insertions(+), 2 deletions(-)

 9ab56ea Delete target line without replacement (broken link)
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Fri Jul 11 02:34:50 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Fri Jul 11 03:36:54 2025 +0000

 gl-test-data/subdir/renamed_target.txt | 1 -
 gl-test-data/test_link2.md             | 2 +-
 gl-test.json                           | 6 +++---
 3 files changed, 4 insertions(+), 5 deletions(-)

 0dd3fd3 Add multi-commit tracking target file
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Sat Jul 12 03:35:51 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Sat Jul 12 04:37:55 2025 +0000

 gl-test-data/multi_link.md    |  3 +++
 gl-test-data/multi_target.txt | 20 ++++++++++++++++++++
 2 files changed, 23 insertions(+)

 f0c12d8 Insert 5 lines at L5 in multi_target
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Sun Jul 13 04:36:52 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Sun Jul 13 05:38:56 2025 +0000

 gl-test-data/multi_target.txt | 5 +++++
 1 file changed, 5 insertions(+)

 c2d8c6a Delete 2 lines at beginning of multi_target
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Mon Jul 14 05:37:53 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Mon Jul 14 06:39:57 2025 +0000

 gl-test-data/multi_target.txt | 2 --
 1 file changed, 2 deletions(-)

 749c6d2 Minor modification to multi_target line
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Tue Jul 15 06:38:54 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Tue Jul 15 07:40:58 2025 +0000

 gl-test-data/multi_target.txt | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)

 820a162 Add absolute path link
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Wed Jul 16 07:39:55 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Wed Jul 16 08:41:59 2025 +0000

 gl-test-data/absolute_link.md | 5 +++++
 1 file changed, 5 insertions(+)

 730388d Add link in Go source file comment
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Thu Jul 17 08:40:56 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Thu Jul 17 09:43:00 2025 +0000

 gl-test-data/code_with_link.go | 16 ++++++++++++++++
 1 file changed, 16 insertions(+)

 c288b86 Test whitespace changes around target
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Fri Jul 18 09:41:57 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Fri Jul 18 10:44:01 2025 +0000

 gl-test-data/multi_target.txt | 5 ++++-
 1 file changed, 4 insertions(+), 1 deletion(-)

 07d7c83 Add large target file for performance testing
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Sat Jul 19 10:42:58 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Sat Jul 19 11:45:02 2025 +0000

 gl-test-data/large_links.md   |   9 ++++
 gl-test-data/large_target.txt | 100 ++++++++++++++++++++++++++++++++++++++++++
 2 files changed, 109 insertions(+)

 88752c0 Modify large file with multiple insertions and deletions
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Sun Jul 20 11:43:59 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Sun Jul 20 12:46:03 2025 +0000

 gl-test-data/large_target.txt | 11 ++++++++---
 1 file changed, 8 insertions(+), 3 deletions(-)

 5a8abcd Add link in file that will be modified in working directory
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Mon Jul 21 12:45:00 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Mon Jul 21 13:47:04 2025 +0000

 gl-test-data/modify_source_link.md    | 1 +
 gl-test-data/modify_source_target.txt | 5 +++++
 2 files changed, 6 insertions(+)

 b325807 Shift target line for modified source file test
    Author: GitRgen Bot <git-rgen-bot@mj41.cz> Tue Jul 22 13:46:01 2025 +0000
    Commit: GitRgen Bot <git-rgen-bot@mj41.cz> Tue Jul 22 14:48:05 2025 +0000

 gl-test-data/modify_source_target.txt |  2 ++
 gl-test.json                          | 34 +++++++++++++++++++++++++++-------
 2 files changed, 29 insertions(+), 7 deletions(-)

Generated repository simple log:
* b325807 (HEAD -> main) Shift target line for modified source file test
* 5a8abcd Add link in file that will be modified in working directory
* 88752c0 Modify large file with multiple insertions and deletions
* 07d7c83 Add large target file for performance testing
* c288b86 Test whitespace changes around target
* 730388d Add link in Go source file comment
* 820a162 Add absolute path link
* 749c6d2 Minor modification to multi_target line
* c2d8c6a Delete 2 lines at beginning of multi_target
* f0c12d8 Insert 5 lines at L5 in multi_target
* 0dd3fd3 Add multi-commit tracking target file
* 9ab56ea Delete target line without replacement (broken link)
* af9d679 Delete target line and add similar line nearby
* 6844f71 Manual link update to reset origin
* 6a02c46 Add nested relative path link using ../
* 1244c7b Add second link file pointing to same target
* 08ecf09 Move target file to subdirectory
* f77a101 Rename target file
* e545fcd Major modification to target line (~60%)
* b8feebd Minor modification to target line (~20%)
* db4a3ae Delete 1 line before target line
* 3756ccf Insert 3 lines after target line
* 6fc4ac2 Insert 2 lines before target line
* 88546a7 Add test link
* 654caff Add test target
* 54949bc Update spec with intro
* a3d7f77 Change docs
* c93c24e Implement v0.0.1
* 8f71e3e Adding simple spec my-tiny-spec.md
* 39ef7fc Repo initialization and simple readme

```
## Script code
```bash
#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR/.."

set -e

echo "Source repository gl-exA-src latest commit:"
cd ../gl-exA-src
git log --oneline -n1
echo

echo "Repository generation:"
cd ../git-rgen-tool
echo "Removing existing gl-exA directory and regenerating with git-rgen..."
echo "Commands:"
echo "  rm -rf ../gl-exA"
echo "  go run ./cmd/rgen --conf ../gl-exA-src/rgen-conf.json"
echo
rm -rf ../gl-exA
go run ./cmd/rgen --conf ../gl-exA-src/rgen-conf.json
echo

echo "Generated repository full log:"
cd ../gl-exA
git log --oneline --decorate --date-order --reverse --stat --format="%n %h %s%n    Author: %an <%ae> %ad%n    Commit: %cn <%ce> %cd"
echo ""

echo "Generated repository simple log:"
git log --oneline --decorate --date-order --graph
echo ""
```

