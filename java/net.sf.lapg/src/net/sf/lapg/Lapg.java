/**
 * Copyright 2002-2010 Evgeny Gryaznov
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package net.sf.lapg;

import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.InputStream;

import net.sf.lapg.api.ProcessingStatus;
import net.sf.lapg.common.FileUtil;
import net.sf.lapg.gen.ConsoleGenerator;
import net.sf.lapg.gen.INotifier;
import net.sf.lapg.gen.LapgOptions;
import net.sf.lapg.gen.ProcessingStatusAdapter;
import net.sf.lapg.parser.LapgTree.TextSource;

/**
 * Main console entry point for Lapg engine.
 */
public class Lapg {

	public static final String VERSION = "1.4.0/java";
	public static final String BUILD = "2010";
	public static final String DEFAULT_FILE = "syntax";

	public static final String HELP_MESSAGE =
		"lapg - Lexer and Parser generator\n"+
		"usage: lapg [OPTIONS]... [inputfile]\n"+
		"       lapg [-h|-v]\n"+
		"\n"+
		"Options:\n"+
		LapgOptions.HELP_OPTIONS+
		"\n"+
		"Operations:\n"+
		"  -h,  --help                    display this help\n"+
		"  -v,  --version                 version information\n"+
		"\n"+
		"Defaults:\n"+
		"  inputfile = "+DEFAULT_FILE+"\n";

	public static final String VERSION_MESSAGE =
		"lapg v" + VERSION + " build " + BUILD + "\n" +
		"Evgeny Gryaznov, 2002-10, egryaznov@gmail.com\n";


	public static void main(String[] args) {
		if (args.length >= 1 && args[0] != null) {
			if (args[0].equals("-h") || args[0].equals("--help")) {
				System.out.println(HELP_MESSAGE);
				return;
			}
			if (args[0].equals("-v") || args[0].equals("--version")) {
				System.out.println(VERSION_MESSAGE);
				return;
			}
		}

		if (args.length == 0 && !new File(DEFAULT_FILE).exists()) {
			System.err.println("lapg: file not found: " + DEFAULT_FILE);
			System.out.println(HELP_MESSAGE);
			System.exit(1);
			return;
		}

		LapgOptions options = LapgOptions.parseArguments(args, System.err);
		if (options == null) {
			System.err.println("Try 'lapg --help' for more information.");
			System.exit(1);
			return;
		}

		InputStream stream;
		if (options.getInput() != null && !options.getInput().startsWith("-")) {
			try {
				stream = new FileInputStream(options.getInput());
			} catch (FileNotFoundException ex) {
				System.err.println("lapg: file not found: " + options.getInput());
				System.exit(1);
				return;
			}
		} else {
			stream = System.in;
		}
		String contents = FileUtil.getFileContents(stream, FileUtil.DEFAULT_ENCODING);
		if(contents == null) {
			System.err.println("lapg: cannot read file: " + options.getInput());
			System.exit(1);
			return;
		}
		TextSource input = new TextSource(options.getInput(), contents.toCharArray(), 1);

		ConsoleGenerator cg = new ConsoleGenerator(options);
		INotifier notifier = cg.createNotifier();
		ProcessingStatus status = new ProcessingStatusAdapter(notifier, options.getDebug());
		boolean success;
		try {
			success = cg.compileGrammar(input, status);
		} finally {
			notifier.dispose();
		}
		if(!success) {
			System.exit(1);
		}
	}
}
