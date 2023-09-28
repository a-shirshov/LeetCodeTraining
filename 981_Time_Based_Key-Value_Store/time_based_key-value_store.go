package timebasedkeyvaluestore

type valueWithTime struct {
    value string
    timestamp int
}

type TimeMap struct {
    hashMap map[string][]valueWithTime
}


func Constructor() TimeMap {
    hashMap := make(map[string][]valueWithTime, 0)
    return TimeMap{
        hashMap: hashMap,
    }
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    this.hashMap[key] = append(this.hashMap[key], valueWithTime{
		value: value,
		timestamp: timestamp,
	})
}


func (this *TimeMap) Get(key string, timestamp int) string {
  values, exist := this.hashMap[key]
	if !exist {
		return ""
	}

	l := 0
	r := len(values) - 1
	saveValue := ""
	for l <= r {
		m := (l+r)/2

		if timestamp == values[m].timestamp {
			return values[m].value
		}

		if timestamp > values[m].timestamp {
			saveValue = values[m].value
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return saveValue
}