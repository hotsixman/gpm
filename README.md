# Go Process Manager (GPM)

PM2의 핵심 기능을 벤치마킹한 Go 기반의 가벼운 프로세스 매니저입니다.

## 🚀 주요 기능 (Key Features)

- **데몬 관리 (Daemonization)**: `gpm init` 명령어를 통해 백그라운드에서 상주하는 관리 프로세스(Daemon)를 실행합니다.
- **자기 복제형 데몬화 (Self-Daemonize)**: 별도의 설정 없이 실행 파일을 재실행하여 터미널 세션과 분리된 독립적인 프로세스를 생성합니다.
- **PID 관리 및 중복 실행 방지**: SQLite 기반의 데이터베이스를 사용하여 현재 실행 중인 데몬의 PID를 관리하고 중복 실행을 방지합니다.
- **플랫폼 지원**: Unix (`Setsid`) 및 Windows (`HideWindow`) 환경에서의 데몬화를 모두 지원합니다.
- **로그 시스템**: 데몬 및 관리 프로세스의 로그를 파일로 기록합니다.

## 📂 프로젝트 구조 (Project Structure)

```text
.
├── main.go             # 엔트리 포인트 (데몬/CLI 분기 처리)
├── module/
│   ├── cli/            # Cobra 기반 CLI 명령어 정의 (init, connect 등)
│   ├── daemon/         # 백그라운드 전환 및 OS별 데몬화 로직
│   ├── database/       # SQLite 기반 PID 및 상태 관리
│   ├── logger/         # 시스템 로그 기록 모듈
│   └── uds/            # Unix Domain Socket 통신 모듈
└── spec.md             # 프로젝트 상세 명세서
```

## 🛠 실행 방법 (Usage)

### 1. 빌드
```bash
go build -o gpm main.go
```

### 2. 데몬 시작 (Initialize)
```bash
./gpm init
```
이 명령어는 부모 프로세스를 즉시 종료하고 백그라운드에서 `gpm gpmd` 프로세스를 실행합니다.

### 3. 상태 확인
현재 실행 중인 데몬 프로세스는 다음과 같이 확인할 수 있습니다.
```bash
ps aux | grep "gpm"
```

## 📜 동작 원리

1. **CLI 실행**: 사용자가 `gpm init`을 호출하면 `daemon.Daemonize()`가 실행됩니다.
2. **프로세스 재생성**: `exec.Command`를 통해 현재 실행 파일에 `GPM_DAEMON_PROCESS=1` 환경변수를 담아 자식 프로세스를 생성합니다.
3. **세션 분리**: `Setsid` 설정을 통해 자식 프로세스를 새로운 세션으로 분리하고 부모 프로세스는 즉시 종료됩니다.
4. **데몬 초기화**: 새 프로세스는 `main.go`에서 환경변수를 확인하여 `daemon.DaemonInit()`으로 진입, DB에 PID를 기록하고 UDS 서버를 준비합니다.

## 🚧 로드맵 (Roadmap)

- [ ] `gpm start <app>`: 새로운 프로세스 등록 및 실행 기능
- [ ] `gpm list`: 현재 관리 중인 프로세스 목록 출력
- [ ] `gpm stop`: 데몬 및 관리 프로세스 안전 종료
- [ ] 실시간 로그 스트리밍 지원 (UDS Broadcast)
