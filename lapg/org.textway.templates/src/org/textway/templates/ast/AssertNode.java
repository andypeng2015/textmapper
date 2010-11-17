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
package org.textway.templates.ast;

import org.textway.templates.api.EvaluationContext;
import org.textway.templates.api.EvaluationException;
import org.textway.templates.api.IEvaluationStrategy;
import org.textway.templates.ast.TemplatesTree.TextSource;

public class AssertNode extends Node {

	private final ExpressionNode expr;

	public AssertNode(ExpressionNode expr, TextSource source, int offset, int endoffset) {
		super(source, offset, endoffset);
		this.expr = expr;
	}

	@Override
	protected void emit(StringBuilder sb, EvaluationContext context, IEvaluationStrategy env) {
		try {
			Object res = env.evaluate(expr, context, true);
			Boolean b = env.toBoolean(res);
			if (!b.booleanValue()) {
				env.fireError(this, "Assertion `" + expr.toString() + "` failed for " + env.getTitle(context.getThisObject()));
			}
		} catch (EvaluationException ex) {
		}
	}
}
