#!/usr/bin/env python3
import json
import shutil
import sys
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
MATRIX = ROOT / "packaging" / "platforms.json"
DIST = ROOT / "dist"


VOICE_FILES = [
    "voices/sherpa/vits-zh-hf-fanchen-C/vits-zh-hf-fanchen-C.onnx",
    "voices/sherpa/vits-zh-hf-fanchen-C/lexicon.txt",
    "voices/sherpa/vits-zh-hf-fanchen-C/tokens.txt",
    "voices/sherpa/vits-zh-hf-fanchen-C/phone.fst",
    "voices/sherpa/vits-zh-hf-fanchen-C/date.fst",
    "voices/sherpa/vits-zh-hf-fanchen-C/number.fst",
    "voices/sherpa/vits-zh-hf-fanchen-C/new_heteronym.fst",
]


def load_targets() -> list[dict]:
    return json.loads(MATRIX.read_text(encoding="utf-8"))["targets"]


def go_available() -> bool:
    bundled = ROOT / "tools" / "go" / "bin" / ("go.exe" if sys.platform.startswith("win") else "go")
    return bundled.exists() or shutil.which("go") is not None


def missing_for_target(target: dict) -> list[str]:
    missing = []
    runtime = ROOT / target["runtime"] / "sherpa-onnx"
    if not runtime.is_dir():
        missing.append(str(runtime.relative_to(ROOT)))
    for item in VOICE_FILES:
        if not (ROOT / item).is_file():
            missing.append(item)
    if target["os"] == "windows":
        arch_label = "x86" if target["arch"] in {"386", "x86"} else target["arch"]
        daemon = DIST / f"wps-tts-daemon-windows-{arch_label}.exe"
        if not daemon.is_file() and not go_available():
            missing.append(str(daemon.relative_to(ROOT)) + " 或可用 Go 工具链")
    elif target["os"] == "linux":
        daemon = DIST / f"wps-tts-daemon-linux-{target['arch']}"
        if not daemon.is_file() and not go_available():
            missing.append(str(daemon.relative_to(ROOT)) + " 或可用 Go 工具链")
    else:
        missing.append(f"unsupported target os: {target['os']}")
    return missing


def main() -> None:
    targets = load_targets()
    all_missing = {}
    for target in targets:
        missing = missing_for_target(target)
        if missing:
            all_missing[target["id"]] = missing
    if all_missing:
        print("platform inputs are incomplete; refusing to build a five-package release.")
        for target_id, items in all_missing.items():
            print(f"[{target_id}]")
            for item in items:
                print(f"  MISSING {item}")
        raise SystemExit(1)
    print("platform_inputs_ok")


if __name__ == "__main__":
    main()
