// Code generated by "callbackgen -type Float64Updater"; DO NOT EDIT.

package indicator

import ()

func (f *Float64Updater) OnUpdate(cb func(v float64)) {
	f.updateCallbacks = append(f.updateCallbacks, cb)
}

func (f *Float64Updater) EmitUpdate(v float64) {
	for _, cb := range f.updateCallbacks {
		cb(v)
	}
}
