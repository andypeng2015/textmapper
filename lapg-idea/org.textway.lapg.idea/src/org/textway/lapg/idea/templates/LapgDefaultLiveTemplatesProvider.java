/**
 * Copyright (c) 2010-2012 Evgeny Gryaznov
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see http://www.gnu.org/licenses/.
 */
package org.textway.lapg.idea.templates;

import com.intellij.codeInsight.template.impl.DefaultLiveTemplatesProvider;
import org.jetbrains.annotations.NonNls;

public class LapgDefaultLiveTemplatesProvider implements DefaultLiveTemplatesProvider {
	private static final
	@NonNls
	String[] DEFAULT_TEMPLATES = new String[]{
			"/liveTemplates/lapg",
	};

	public String[] getDefaultLiveTemplateFiles() {
		return DEFAULT_TEMPLATES;
	}

	public String[] getHiddenLiveTemplateFiles() {
		return null;
	}
}
