#!/usr/bin/env python3
import argparse
import filecmp
import shutil
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
ADDIN = ROOT / "addin"
WEB = ROOT / "daemon" / "cmd" / "wps-tts-daemon" / "web"


def compare_dirs(left: Path, right: Path) -> list[str]:
    if not left.exists():
        return [f"missing source directory: {left}"]
    if not right.exists():
        return [f"missing embedded directory: {right}"]

    diffs: list[str] = []
    cmp = filecmp.dircmp(left, right)
    for name in cmp.left_only:
        diffs.append(f"missing in embedded web: {left / name}")
    for name in cmp.right_only:
        diffs.append(f"extra in embedded web: {right / name}")
    for name in cmp.diff_files:
        diffs.append(f"content differs: {left / name}")
    for sub in cmp.common_dirs:
        diffs.extend(compare_dirs(left / sub, right / sub))
    return diffs


def sync() -> None:
    if WEB.exists():
        shutil.rmtree(WEB)
    WEB.mkdir(parents=True, exist_ok=True)
    shutil.copytree(ADDIN, WEB, dirs_exist_ok=True)


def main() -> None:
    parser = argparse.ArgumentParser(description="sync WPS add-in assets into Go embedded web directory")
    parser.add_argument("--check", action="store_true", help="only verify that embedded web matches addin")
    args = parser.parse_args()

    if args.check:
        diffs = compare_dirs(ADDIN, WEB)
        if diffs:
            raise SystemExit("embedded web assets are not synchronized:\n" + "\n".join(diffs))
        print("embedded web assets are synchronized")
        return

    sync()
    print(f"synced {ADDIN} -> {WEB}")


if __name__ == "__main__":
    main()
