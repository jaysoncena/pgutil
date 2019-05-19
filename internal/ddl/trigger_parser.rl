package ddl

// Code generated by go generate; DO NOT EDIT.

import (
	"bytes"
)

%%{
	machine trigger;
	write data;
}%%

func ParseTrigger(data string) (*Trigger, error) {
	trigger := &Trigger{}

	buffer := &bytes.Buffer{}
	cs, p, pe := 0, 0, len(data)
	%%{

	action addToBuffer       {
		buffer.WriteByte(fc)
	}

	action addColumn {
		trigger.addColumn(buffer.String())
		buffer.Reset()
	}
	action setFunction {
		trigger.setFunction(buffer.String())
		buffer.Reset()
	}
	action setName {
		trigger.setName(buffer.String())
		buffer.Reset()
	}
	action setTable {
		trigger.setTable(buffer.String())
		buffer.Reset()
	}
	action setWhen {
		trigger.setWhen(buffer.String())
		buffer.Reset()
	}

	action matchConstraint        { trigger.Constraint = true }
	action matchDeferrable        { trigger.Deferrable = true }
	action matchInitiallyDeferred { trigger.InitiallyDeferred = true }
	action matchForEachRow        { trigger.ForEachRow = true }

	action matchDelete   { trigger.Delete = true }
	action matchInsert   { trigger.Insert = true }
	action matchTruncate { trigger.Truncate = true }
	action matchUpdate   { trigger.Update = true }

	ws = space+;
	ident = [a-zA-Z][_a-zA-Z0-9]*;

	main := space*
		'CREATE'i ws ( 'CONSTRAINT'i @ matchConstraint ws )?  'TRIGGER'i
		ws ( ident $ addToBuffer % setName )
		ws ( ('BEFORE'i | 'AFTER'i | ('INSTEAD'i ws 'OF'i)) $ addToBuffer % setWhen )
		ws ( 'DELETE'i @ matchDelete | 'INSERT'i @ matchInsert | 'TRUNCATE'i @ matchTruncate | 'UPDATE'i @ matchUpdate )
		( ws 'OR'i ws
		   ( 'DELETE'i @ matchDelete | 'INSERT'i @ matchInsert | 'TRUNCATE'i @ matchTruncate | 'UPDATE'i @ matchUpdate )
		)*
		ws 'ON'i ws ( ident $ addToBuffer % setTable )
		( ws ( 'NOT'i ws 'DEFERRABLE'i )
		| ws ( 'DEFERRABLE'i ws @ matchDeferrable )?
		   ( 'INITIALLY'i ws 'IMMEDIATE'i | 'INITIALLY'i ws 'DEFERRED'i @ matchInitiallyDeferred )?
		)?
		( ws 'FOR'i ws ('EACH'i ws)? ('ROW'i @ matchForEachRow | 'STATEMENT'i) )?
		ws 'EXECUTE'i ws ('FUNCTION'i | 'PROCEDURE'i) ws ( ident $ addToBuffer % setFunction ) '()'
		space*
		;

	write init;
	write exec;
	}%%

	if cs < trigger_first_final {
		return nil, &parseError{
		    cs: cs,
		    data: data,
		}
	}

	return trigger, nil
}
