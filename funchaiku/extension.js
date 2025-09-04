// The module 'vscode' contains the VS Code extensibility API
// Import the module and reference it with the alias vscode in your code below
const vscode = require('vscode');
const { LanguageClient } = require('vscode-languageclient/node');

let client;

// This method is called when your extension is activated
// Your extension is activated the very first time the command is executed

/**
 * @param {vscode.ExtensionContext} context
 */
function activate(context) {

	// Use the console to output diagnostic information (console.log) and errors (console.error)
	// This line of code will only be executed once when your extension is activated
	console.log('Congratulations, your extension "funchaiku" is now active!');


	const serverCommand = 'lsp-server';

	const serverOptions = {
		command: serverCommand,
		args: []
	};

	const clientOptions = {
		documentSelector: [
			{ scheme: 'file', language: 'go' },
			{ scheme: 'file', language: 'javascript' },
			{ scheme: 'file', language: 'typescript' },
		],
		synchronize: {
			fileEvents: vscode.workspace.createFileSystemWatcher('**/*')
		},
	};

	client = new LanguageClient(
		'funchaiku',
		'Funchaiku Language Server',
		serverOptions,
		clientOptions
	);

	client.start();

	context.subscriptions.push(client);

	// Command to manually trigger haiku generation
	const haikuCommand = vscode.commands.registerCommand(
		'funchaiku.generateHaiku',
		() => {
			const editor = vscode.window.activeTextEditor;
			if (editor) {
				const position = editor.selection.active;
				vscode.commands.executeCommand(
					'vscode.executeHoverProvider',
					editor.document.uri,
					position
				)
			}
		}
	);
	context.subscriptions.push(haikuCommand);
}

// This method is called when your extension is deactivated
function deactivate() {
	if (!client) {
		return;
	}
	return client.stop();
}

module.exports = {
	activate,
	deactivate
}
