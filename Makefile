fulllint: fmtcheck importscheck vetcheck gocriticcheck

test:
	@sh -c "'$(CURDIR)/scripts/test.sh'"

vetcheck:
	@sh -c "'$(CURDIR)/scripts/vet.sh'"

gocriticcheck:
	@sh -c "'$(CURDIR)/scripts/gocritic.sh'"

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/format.sh'"

importscheck:
	@sh -c "'$(CURDIR)/scripts/import.sh'"

.PHONY: fulllint test vetcheck gocriticcheck fmtcheck importscheck
