# Changelog
本專案所有的顯著修改都將記錄在此文件中。

文件依照 [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) 所描述的格式撰寫，
版本號碼格式依照 [Semantic Versioning](https://semver.org/spec/v2.0.0.html)。

## [Unreleased]

## [0.14.0] - 2021-10-24

### Added
- 公車V2版之進階資料服務上架。

### Changed
- 航運API V3版資料服務更新。
- 軌道V2版資料服務上架與修復。

## [0.13.0] - 2021-09-20

### Added
- 捷運V2版資料服務上架。

### Changed
- 公總及代管縣市市區公車資料服務內容更新。
- 自行車資料服務內容更新。
- 高雄捷運、高雄輕軌資料服務預計資料結構統一調整。
- 捷運V2版淡海輕軌時刻表資料服務更新。
- 航空及觀光 API 新增 status 299 回應。

## [0.12.0] - 2021-07-31

### Added
- 捷運、公共自行車V2版資料服務上架。

### Changed
- 阿里山小火車、公共自行車、公車資料服務更新。
- 自行車及軌道 API 新增 status 299 回應。

### Fixed
- 桃園公車資料服務內容更新。

## [0.11.0] - 2021-07-10

### Added
- 高雄市公車新資料服務上架
- 平臺API全面提供HTTP 304狀態碼查詢服務
- 高雄輕軌V2版資料服務上架
- 軌道V3版資料服務上架
- 公車V2版資料服務上架

### Removed
- 彰化縣公共自行車服務下架

## [0.10.0] - 2021-04-24

### Added
- 臺中市公車新資料服務上架
- 公車新資料服務上架

### Changed
- 公車站位資料服務內容更新
- 公車 V2 API 函數名稱修改

### Workaround
- 修正公車 V2 多處 Direction 型別錯誤

## [0.9.0] - 2021-02-02

### Added
- 公路客運線型資料服務上架
- 公車、軌道與自行車新資料服務上架
- 航運V3版船舶即時位置資料服務上架

### Changed
- 軌道資料服務更新
- 高鐵動態即時剩餘位資料服務預計資料結構調整
- 航空資料服務更新
- 台鐵、高鐵資料服務預計資料結構調整
- 公車預估到站(N1)服務內容更新
- 高鐵動態即時剩餘位資料服務資料結構調整
- 淡海輕軌資料服務更新

## [0.8.0] - 2020-11-02

### Changed
- 公車資料服務更新。
- 基礎資料及自行車服務更新。

## [0.7.0] - 2020-09-30

### Added
- 高雄市公共自行車服務。
- 公車服務更新。
- 高鐵對號座即時剩餘位資料服務(加值型列車起迄段OD角度)上架。

### Changed
- 北捷V2版車站基本資料服務更新。
- 軌道、自行車及觀光服務更新。
- 高鐵V2每日時刻表資料結構已調整。

## [0.6.0] - 2020-08-27

### Added
- 金門縣公車及自行車服務。
- 北捷貓纜V2版本資料服務上架。
- 高雄輕軌V2版資料服務上架。

### Changed
- 高鐵V2每日時刻表資料結構調整。
- 高鐵定期時刻表資料修正。

### Removed
- 高雄市公共自行車服務。

## [0.5.0] - 2020-06-18

### Added
- 金門公車資料服務。
- 高雄捷運、高雄輕軌V2版News、Alert資料服務上架。
- 淡海輕軌V2版資料服務上架。
- 捷運環狀線V2版資料服務上架。
- 軌道資料服務更新。

### Fixed
- 觀光台灣好行預估到站資料服務。
- 航空資料服務更新。

## [0.4.0] - 2020-01-23

### Added
- 公車 V2 新增逐筆更新資料。
- 軌道 V2 新增桃園捷運營運最新消息資料。
- 軌道 V2 新增桃園捷運營運通阻資料。

### Changed
- 公車 V2 改版。
- 軌道 V2 台鐵車站資料新增必填欄位 `StationUID`。

## [0.3.0] - 2019-12-21

### Added
- 官方新增公車 V3 API `/v3/Bus/Schedule/City/{City}`。
- 官方新增軌道 V3 API `/v3/Rail/TRA/Operator`、`/v3/Rail/TRA/LineNetwork`。

### Changed
- 官方新增 health 參數，查詢 API 服務健康狀態。
- 改進 CI 流程。
- 改進文件。

### Fixed
- 修正 check-oas-spec.sh 檢查 SDK 更新的錯誤。

## [0.2.0] - 2019-09-28

### Changed
- 官方新增公車 V2 API 預估到站時間，目前僅桃園市、臺中市、高雄市提供。
- 官方修改公車 V3 API operationId。
- 修改 struct filed `AppId` 為 `AppID`，以符合 golint 規則。
- 改進 CI 流程，預設使用 go1.13 測試。
- 改用最新的 go-swagger static binary 來驗證定義檔。

### Fixed
- 修正 rail v2 測試。

## [0.1.0] - 2019-08-04

### Added
- 支援公車 V3 API（僅提供台南市公車資訊）。

### Changed
- 官方修改 API 參數 `$top` 的型別為 `integer`。
- 官方修改軌道 V2 API `/v2/Rail/Metro/Network/{Operator}`，新增欲查詢縣市（`Operator`）參數。
- 官方修改觀光 API tag `taiwan_trip_bus` 為 `tourism`，重新產生用戶端程式庫。

### Fixed
- 配合官方修正 `enum` 用法錯誤，移除 `make spec` 內臨時的解決方案。

## [0.0.1] - 2019-07-20

### Added
- 首次發行，提供基本、航空、公車、軌道、自行車與觀光等 PTX API 用戶端程式庫。

[Unreleased]: https://github.com/minchao/go-ptx/compare/v0.14.0...HEAD
[0.14.0]: https://github.com/minchao/go-ptx/compare/v0.13.0...v0.14.0
[0.13.0]: https://github.com/minchao/go-ptx/compare/v0.12.0...v0.13.0
[0.12.0]: https://github.com/minchao/go-ptx/compare/v0.11.0...v0.12.0
[0.11.0]: https://github.com/minchao/go-ptx/compare/v0.10.0...v0.11.0
[0.10.0]: https://github.com/minchao/go-ptx/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/minchao/go-ptx/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/minchao/go-ptx/compare/v0.7.0...v0.8.0
[0.7.0]: https://github.com/minchao/go-ptx/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/minchao/go-ptx/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/minchao/go-ptx/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/minchao/go-ptx/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/minchao/go-ptx/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/minchao/go-ptx/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/minchao/go-ptx/compare/v0.0.1...v0.1.0
[0.0.1]: https://github.com/minchao/go-ptx/commits/v0.0.1
