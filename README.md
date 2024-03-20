# Creating an LSP From scratch 

## Installing the project 
### Clone the repo 

```bash
    git clone https://github.com/Sarath191181208/<repo_url>
```
if you have not installed go [install](https://go.dev/doc/install).
build the project using go.

```bash 
go build main.go
```


### triggering LSP in nvim
```lua
local client = vim.lsp.start_client{
  name = "test_lsp",
  -- Make sure to enter the correct path
  cmd = { "~/Projects/lsp_from_scratch/main" },
  on_attach = require("plugins.configs.lspconfig").on_attach
}

if not client then
  vim.notify("Hey, the client isn't configured properly")
  return
end

vim.api.nvim_create_autocmd('FileType', {
  pattern = "markdown",
  callback = function ()
    vim.lsp.buf_attach_client(0, client)
  end
})
```
