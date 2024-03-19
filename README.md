
### triggering LSP in nvim
```lua
local client = vim.lsp.start_client{
  name = "test_lsp",
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
