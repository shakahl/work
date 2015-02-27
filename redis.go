package work

func redisNamespacePrefix(namespace string) string {
	l := len(namespace)
	if (l > 0) && (namespace[l-1] != ':') {
		namespace = namespace + ":"
	}
	return namespace
}

// returns "<namespace>:jobs:"
// so that we can just append the job name and be good to go
func redisKeyJobsPrefix(namespace string) string {
	return redisNamespacePrefix(namespace) + "jobs:"
}

func redisKeyJobs(namespace, jobName string) string {
	return redisKeyJobsPrefix(namespace) + jobName
}

func redisKeyJobsInProgress(namespace, jobName string) string {
	return redisKeyJobs(namespace, jobName) + ":inprogress"
}

func redisKeyRetry(namespace string) string {
	return redisNamespacePrefix(namespace) + "retry"
}

func redisKeyDead(namespace string) string {
	return redisNamespacePrefix(namespace) + "dead"
}

var redisLuaRpoplpushMultiCmd = `
local res
local keylen = #KEYS
for i=1,keylen,2 do
  res = redis.call('rpop', KEYS[i])
  if res then
    redis.call('lpush', KEYS[i+1], res)
    return {res, KEYS[i], KEYS[i+1]}
  end
end
return nil`