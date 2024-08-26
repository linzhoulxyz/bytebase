# 自编译构建步骤

## 步骤
1. 同步github仓库:
   1. https://github.com/xxx/bytebase
2. 本地同步
   1. git pull
3. 拉取官方仓库最新分支: 2.22.2
   1. git fetch upstream
4. 使用官方仓库最新分支创建新分支
   1. git checkout -b release/2.22.2 upstream/release/2.22.2
5. 将新创建的分支推送到github
   1. git push origin release/2.22.2
6. 找到上个分支，未提交的commit id
   1. git checkout release/2.22.0
   2. git log -3
7. 切换回最新的分支，并cherry-pick提交
      1. git checkout release/2.22.2
      2. git cherry-pick f9c1dc63f3b13568da9bcea73067aa2a24d70ec7..1bc194aa637aa861bea990d6434a6498d5b0f398
8. 如果有冲突，解决冲突
9. 将刚才的提交生成patch
      1. git format-patch -1 commit_hash
      2. git format-patch commit_hash1..commit_mdash2
10. 将生成的patch复制到编译服务器，并应用patch
      1. git apply patch_file.patch