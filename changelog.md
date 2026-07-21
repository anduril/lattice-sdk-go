# Changelog

## [v5.0.0] - 2026-07-21
### Breaking Changes
- **`Entity` field bitmask ordering** — the new `Kinematics` field was inserted after `LocationUncertainty`, shifting the explicit-field bit positions of `GeoShape`, `GeoDetails`, and all subsequent fields; regenerate and re-serialize any persisted `explicitFields` masks.

### Added
- **`Altitude`** and related types (`AltitudeAboveWgs84Ellipsoid`, `AltitudeAboveSeaFloor`, `AltitudeBelowSeaSurface`, `AltitudeAboveStandardDatumPlanePressure`, `AltitudeAboveMeanSeaLevelPressure`, `AltitudeAboveMeanSeaLevelEgm96`, `AltitudeAboveGroundLevel`) for expressing altitude in multiple reference frames.
- **`AltitudeProvenance`** and **`AltitudeProvenanceSourceType`** enum for tracking how an altitude measurement was derived.
- **`Kinematics`**, **`KinematicsGeodetic`**, and **`KinematicsGeocentric`** for higher-fidelity entity motion data, plus supporting **`LocationGeodetic`**, **`LocationGeocentricEcef`**, **`TMat3`**, and **`Vec3`** types.
- **`Entity.Kinematics`** field (with `GetKinematics`/`SetKinematics`) as an alternative to `Location`/`LocationUncertainty`.

## [4.19.0] - 2026-07-16

## [v4.18.1] - 2026-07-14

