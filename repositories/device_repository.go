package repositories

type DeviceRepository struct{}

func NewDeviceRepository() *DeviceRepository {
	return &DeviceRepository{}
}

func (dr *DeviceRepository) FindDevice(customerID, globalID, imei, androidID, advID string) string {
	// Mock deviceID selection
	if globalID != "" {
		return "device_001"
	}
	return ""
}
