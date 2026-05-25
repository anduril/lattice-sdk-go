#!/usr/bin/env python3

import re
from pathlib import Path

FILES = [
    "entities/raw_client.go",
    "entities/client.go",
    "objects/raw_client.go",
    "objects/client.go",
    "tasks/raw_client.go",
    "tasks/client.go",
    "oauth/raw_client.go",
]

PATTERN = re.compile(
    r"(?m)^(\t*)headers := internal\.MergeHeaders\(\n"
    r"\1\t([rc])\.options\.ToHeader\(\),\n"
    r"\1\toptions\.ToHeader\(\),\n"
    r"\1\)"
)


def replacement(match: re.Match) -> str:
    indent = match.group(1)
    receiver = match.group(2)

    return (
        f"{indent}headers, err := internal.MergeRequestHeaders({receiver}.options, options)\n"
        f"{indent}if err != nil {{\n"
        f"{indent}\treturn nil, err\n"
        f"{indent}}}"
    )


def main() -> None:
    changed = []

    for filename in FILES:
        path = Path(filename)
        original = path.read_text()
        updated, count = PATTERN.subn(replacement, original)

        if count == 0:
            raise SystemExit(f"no header call sites patched in {filename}")

        path.write_text(updated)
        changed.append((filename, count))

    for filename, count in changed:
        print(f"patched {count} call site(s) in {filename}")


if __name__ == "__main__":
    main()
