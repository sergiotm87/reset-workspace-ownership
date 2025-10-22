# reset-workspace-ownership
Github Action to ensure correct workspace ownership. Ideal for self-hosted runners

## Usage

```yaml
- name: Reset workspace ownership
  uses: ./
  with:
    path: 'path/to/your/workspace'
    uid: '1001'
```

## Inputs

| Name   | Description                               | Required |
|--------|-------------------------------------------|----------|
| `path` | The path to the directory to reset ownership on. | `true`   |
| `uid`  | The user ID to set as the owner.          | `true`   |
