package net.sf.lapg.gen.options.ast;

import net.sf.lapg.gen.options.OptdefTree.TextSource;

public interface IAstOptNode {
	int getOffset();
	int getEndOffset();
	TextSource getInput();
	//void accept(Visitor v);
}
