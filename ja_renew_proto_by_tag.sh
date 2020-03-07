#!/bin/bash

echo "[Jean]: On va remove .gitmodule  .git/modules/proto pour toi "
echo "[Jean]: Ensuite , creer nouvelle avec the tag version of proto "
echo "--------------------------------------------"
git ls-remote --tags git@gitlab.com:Johnny_Wick/ja_gitlab_grpc_proto.git
echo "------------------------------------------"
echo "Ta sure ?"
echo "Enter tag name: v1.0.2"
read tag_name


rm -rf proto

echo "------------------------------"
echo " rm -rf .git/modules/proto/ "
echo " rm .gitmodules             "
echo "------------------------------"
rm -rf .git/modules/proto/
rm .gitmodules




echo "[Jean]: Mnt. je suppliser tout le files"
echo "       tu dois commit back to gitlab first"
echo "Enter commit message = 'clean'"
echo "Are you ready ? "
read nothing_var
source ja_git_push_back_v2.sh


echo
echo
echo
echo "########################################################################"

echo "[Jean]: Cloning... "
echo "Example: "
echo "git submodule add git@gitlab.com:Johnny_Wick/ja_gitlab_grpc_proto.git proto"
git submodule add -b $tag_name git@gitlab.com:Johnny_Wick/ja_gitlab_grpc_proto.git proto
