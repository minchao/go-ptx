# Changelog
本專案所有的顯著修改都將記錄在此文件中。

文件依照 [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) 所描述的格式撰寫，
版本號碼格式依照 [Semantic Versioning](https://semver.org/spec/v2.0.0.html)。

## [Unreleased]

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

[Unreleased]: https://github.com/minchao/go-ptx/compare/v0.1.0...HEAD
[0.1.0]: https://github.com/minchao/go-ptx/compare/v0.0.1...v0.1.0
[0.0.1]: https://github.com/minchao/go-ptx/commits/v0.0.1
