## [UNRELEASED] go-datareader 1.0.0

### Breaking changes
- Add `Tiingo` daily data reader ([#2](https://github.com/AleksanderWWW/go-datareader/pull/2))
- Use config structs to initialize readers ([#5](https://github.com/AleksanderWWW/go-datareader/pull/5))

### Changes
- Repo cleanup ([#1](https://github.com/AleksanderWWW/go-datareader/pull/1))
- Use config struct to initialize `Tiingo` reader ([#3](https://github.com/AleksanderWWW/go-datareader/pull/3))
- Remove logging to file ([#6](https://github.com/AleksanderWWW/go-datareader/pull/6))
- Move base urls to a separate `constants.go` file ([#8](https://github.com/AleksanderWWW/go-datareader/pull/8))

### Fixes
- Handle empty dataframe list in `concatDataframes` ([#9](https://github.com/AleksanderWWW/go-datareader/pull/9))
