ioutil.ReadAll -> io.ReadAll
ioutil.ReadFile -> os.ReadFile
ioutil.ReadDir -> os.ReadDir
// others
ioutil.NopCloser -> io.NopCloser
ioutil.TempDir -> os.MkdirTemp
ioutil.TempFile -> os.CreateTemp
ioutil.WriteFile -> os.WriteFile
