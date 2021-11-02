#!/bin/bash
set -e

# default dev credentials
export PGHOST=localhost
export PGDATABASE=explorer
export PGUSER=bdjuno
export PGPASSWORD=foobarbaz

# setup all the tables
psql -f database/schema/00-cosmos.sql
psql -f database/schema/01-auth.sql
psql -f database/schema/02-bank.sql
psql -f database/schema/03-staking.sql
psql -f database/schema/04-consensus.sql
psql -f database/schema/05-mint.sql
psql -f database/schema/06-distribution.sql
psql -f database/schema/07-pricefeed.sql
psql -f database/schema/08-gov.sql
psql -f database/schema/09-modules.sql
psql -f database/schema/10-slashing.sql
psql -f database/schema/11-history.sql
