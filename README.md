# git-stats

This is a tool used to fetch info and stats about git contributors of a repo.

## build

```
go build -o git-stats .
chmod +x git-stats
```

## usage

```
export TOKEN=$GIT_TOKEN
git-stats -o <github-org> -r <github-repo> -t $TOKEN
```

The tool should then output some info that looks like the following:

```
LOGIN			        CONTRIBUTIONS	    EMAIL				      COMPANY					ORGS
<github-login>		<commit-#>		    <github-email>		<company> 			<github-orgs>
.
.
.
```
