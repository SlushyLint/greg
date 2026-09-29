" =========================
" ~/.vimrc
"
" =========================

" --- Basic ---
" --- ALE ---
let g:go_highlight_fields = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_function_parameters = 1
let g:go_highlight_functions = 1
let g:go_highlight_operators = 1
let g:go_highlight_types = 1
syntax on

set signcolumn=no
set number
set autoindent
set smartindent
set expandtab
set tabstop=4
set shiftwidth=4
set softtabstop=4
set noshowmode
" --- Syntax highlighting ---
syntax enable
set background=dark
set termguicolors
" --- Cursor ---
set guicursor=a:block
highlight Cursor guibg=Red guifg=Black
" --- jj = Escape ---
inoremap jj <Esc>

" --- ; = : in Normal mode ---
nnoremap ; :

" --- Persistent undo ---
set undofile
set undodir=~/.vim/undo
set history=1000

" --- Persistent command/search history ---
set viminfo='1000,<1000,s100,h

" --- Make sure undo directory exists ---
if !isdirectory(expand('~/.vim/undo'))
    call mkdir(expand('~/.vim/undo'), 'p')
endif

" --- Search highlighting ---
set hlsearch
set incsearch

" --- Better indentation ---
filetype plugin indent on

" --- Show matching brackets ---
set showmatch

" --- Don't wrap code ---
set nowrap
" Enable line highlighting and relative numbering
set number

" Customize the current line number color
" guifg changes text color, guibg changes the number's background color
highlight CursorLineNr ctermfg=Yellow gui=bold


" Error line highlighting
nnoremap <C-L> :nohlsearch<CR>

" For the currently selected/active search term
hi CurSearch ctermbg=Green ctermfg=Black guibg=Green guifg=Black
set timeout
set timeoutlen=150
set makeprg=go\ build\ .
set shellcmdflag=-c
autocmd BufRead,BufNewFile *.vala set filetype=c

" 1. --- Plugin Management (Example using vim-plug) ---
let g:lsp_semantic_enabled = 1
let g:lsp_semantic_delay = 0

call plug#begin('~/.vim/plugged')
Plug 'nordtheme/vim'
Plug 'dense-analysis/ale'
Plug 'luochen1990/rainbow'
Plug 'jiangmiao/auto-pairs'
Plug 'prabirshrestha/vim-lsp'
call plug#end()

au User lsp_setup call lsp#register_server({
\   'name': 'gopls',
\   'cmd': ['gopls'],
\   'allowlist': ['go'],
\})
" 2. --- Nord Color Scheme Setup ---
syntax enable
colorscheme nord

" Match Go syntax colors to Nord Dark Pro in VS Code.
highlight goStatement guifg=#81A1C1 ctermfg=Blue
highlight goConditional guifg=#81A1C1 ctermfg=Blue
highlight goRepeat guifg=#81A1C1 ctermfg=Blue
highlight goLabel guifg=#81A1C1 ctermfg=Blue
highlight goDeclaration guifg=#81A1C1 ctermfg=Blue
highlight goTypeDecl guifg=#81A1C1 ctermfg=Blue
highlight goFunction guifg=#88C0D0 ctermfg=Cyan
highlight goFunctionCall guifg=#88C0D0 ctermfg=Cyan
highlight goBuiltins guifg=#88C0D0 ctermfg=Cyan
highlight goField guifg=#6F94B8 ctermfg=Blue
highlight goParamName guifg=#6F94B8 ctermfg=Blue
highlight goType guifg=#8FBCBB ctermfg=LightCyan
highlight goTypeName guifg=#8FBCBB ctermfg=LightCyan
highlight goTypeConstructor guifg=#8FBCBB ctermfg=LightCyan
highlight goString guifg=#A3BE8C ctermfg=Green
highlight goRawString guifg=#A3BE8C ctermfg=Green
highlight goComment guifg=#596477 ctermfg=DarkGray
highlight Delimiter guifg=#C4CDDD ctermfg=LightCyan
highlight goOperator guifg=#81A1C1 ctermfg=Blue

" Keep inactive line numbers dimmer than comments; mark the active row clearly.
set cursorline
highlight LineNr guifg=#3B4252 ctermfg=Black
highlight CursorLine guibg=#2C313D ctermbg=8
highlight CursorLineNr guifg=#ABB2BF ctermfg=White

" 3. --- Keep Terminal Background Color Unchanged ---
" This overrides Nord's default background, forcing it to be transparent/none.
" Must be placed AFTER the 'colorscheme nord' command.
function! NormaliseBackground()
    highlight Normal ctermbg=NONE guibg=NONE
    highlight NonText ctermbg=NONE guibg=NONE
    highlight LineNr ctermbg=NONE guibg=NONE
    highlight SignColumn ctermbg=NONE guibg=NONE
    highlight EndOfBuffer ctermbg=NONE guibg=NONE
endfunction

autocmd ColorScheme nord call NormaliseBackground()
" Run it immediately for the current session
call NormaliseBackground()
" ALE — Go
let g:ale_linters = {
\   'go': ['gopls'],
\}

let g:ale_fixers = {
\   'go': ['gofmt', 'goimports'],
\}

let g:ale_fix_on_save = 1

let g:ale_lint_on_text_changed = 'always'
let g:ale_lint_on_insert_leave = 1
let g:ale_lint_on_enter = 1
let g:ale_lint_delay = 150

" Highlight ONLY the offending source text
let g:ale_set_highlights = 1

highlight ALEError guifg=#000000 guibg=#88C0D0
highlight ALEWarning guifg=#000000 guibg=#88C0D0

highlight ALEError ctermfg=Black ctermbg=Cyan
highlight ALEWarning ctermfg=Black ctermbg=Cyan
" 1. Enable the rainbow plugin globally
let g:rainbow_active = 1

" 2. Force rainbow colors to match the Nord color palette
"  let g:rainbow_conf = {
"  \   'guifgs': ['#8FBCBB', '#88C0D0', '#81A1C1', '#5E81AC', '#A3BE8C', '#EBCB8B'],
"  \   'ctermfgs': ['cyan', 'lightcyan', 'blue', 'darkblue', 'green', 'yellow'],
"  \   'parentheses': ['start=/(/ end=/)/ fold', 'start=/\[/ end=/\]/ fold', 'start=/{/ end=/}/ fold'],
"  \}
"  let g:rainbow_conf = {
"  \   'guifgs': ['#8FBCBB', '#88C0D0', '#81A1C1', '#5E81AC', '#A3BE8C', '#EBCB8B'],
"  \   'ctermfgs': ['cyan', 'lightcyan', 'blue', 'darkblue', 'green', 'yellow'],
"  \   'parentheses': ['start=/(/ end=/)/ fold', 'start=/\[/ end=/\]/ fold', 'start=/{/ end=/}/ fold'],
"  \}

let g:rainbow_conf = {
\   'guifgs': ['#81A1C1', '#8FBCBB', '#A3BE8C', '#EBCB8B', '#B48EAD', '#88C0D0'],
\   'ctermfgs': ['Blue', 'Cyan', 'Green', 'Yellow', 'Magenta', 'LightCyan'],
\   'parentheses': ['start=/(/ end=/)/ fold', 'start=/\[/ end=/\]/ fold', 'start=/{/ end=/}/ fold'],
\}


" 5. Clean up auto-pairs conflict by forcing a syntax refresh on file load
autocmd BufWinEnter,Syntax * RainbowToggleOn





set viminfo='10

augroup remember_cursor_position
    autocmd!
    autocmd BufReadPost * if line("'\"") > 0 && line("'\"") <= line("$") | execute "normal! g`\"" | endif
augroup END
highlight LspSemanticVariable guifg=#6F94B8 ctermfg=Blue
highlight LspSemanticProperty guifg=#6F94B8 ctermfg=Blue
highlight LspSemanticParameter guifg=#6F94B8 ctermfg=Blue
highlight LspSemanticFunction guifg=#88C0D0 ctermfg=Cyan
highlight LspSemanticMethod guifg=#88C0D0 ctermfg=Cyan
highlight LspSemanticType guifg=#8FBCBB ctermfg=LightCyan
highlight LspSemanticString guifg=#A3BE8C ctermfg=Green
highlight LspSemanticNumber guifg=#B48EAD ctermfg=Magenta
let g:lsp_document_highlight_enabled = 1
let g:lsp_document_highlight_delay = 0
