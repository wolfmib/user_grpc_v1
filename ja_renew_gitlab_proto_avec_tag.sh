#!/bin/bash

echo "[Jean]: On va remove .gitmodule  .git/modules/proto pour toi "
echo "[Jean]: Ensuite , creer nouvelle latest version of proto "
echo "Ta sure ?"
echo "Enter ...."
read nothing_var

echo "[Jean]: C'est program! on va utiliser la specific tag ou branch"
echo "Type the tag:"
echo "Par Example: v1.0.2, Tu va demander la engineer ! qui est la suitable tag pour toi maintenant.."
echo 
read tag_cmd



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
git submodule add git@gitlab.com:Johnny_Wick/ja_gitlab_grpc_proto.git proto


# [Jean]: move the submodule to a particular tag:
echo "Pulling the tag in sub-module , proto ...."
echo 
cd proto
git checkout $tag_cmd
echo "Checking..."
echo "---------------"
git status
echo "---------------"
cd ..
echo "[Jean]: Je vais utiliser the ja_init_git_push_v2 pour updated la proto's gitlab version "
echo "Cava ? "
echo 
read nothing_var

source ja_git_push_back_v2.sh


