package mydictionary

import "github.com/zzc-tongji/rtoa"

// dictionart list
type collectionListSlice []collectionStruct

// read all collection from .xlsx file
func (collectionList *collectionListSlice) read(setting *settingStruct) (err error) {
	var (
		collection collectionStruct
		str        string
	)
	// read collection
	for i := 0; i < len(setting.Collection); i++ {
		collection.name = setting.Collection[i].Name
		collection.readable = setting.Collection[i].Readable
		collection.writable = setting.Collection[i].Writable
		collection.onlineSource = setting.Collection[i].OnlineSource
		str, err = rtoa.Convert(setting.Collection[i].FilePath, "")
		if err != nil {
			return
		}
		err = collection.read(str)
		if err != nil {
			return
		}
		*collectionList = append(*collectionList, collection)
	}
	return
}

// write all collection to .xlsx file
func (collectionList *collectionListSlice) write() (success bool, information string) {
	var (
		err  error
		temp string
	)
	success = true
	for i := 0; i < len(*collectionList); i++ {
		temp, err = (*collectionList)[i].write()
		if err != nil {
			temp = err.Error() + "\n\n"
			success = false
		}
		information += temp
	}
	return
}
