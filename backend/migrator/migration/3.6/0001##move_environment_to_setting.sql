ALTER TABLE instance DROP CONSTRAINT IF EXISTS instance_environment_fkey;

ALTER TABLE db DROP CONSTRAINT IF EXISTS db_environment_fkey;

ALTER TABLE stage DROP CONSTRAINT IF EXISTS stage_environment_fkey;


INSERT INTO setting (
    name,
    value
) SELECT 
    'bb.workspace.environment',
    jsonb_build_object(
        'environments',
        jsonb_agg(environments.v)
    )
    FROM (
        SELECT 
            jsonb_build_object('title', name) ||
            jsonb_build_object('id', resource_id) ||
            CASE WHEN policy.payload->>'environmentTier' = 'PROTECTED' THEN jsonb_build_object('tags', jsonb_build_object('protected', 'protected'))
                ELSE '{"tags": {}}' 
            END 
            AS v
            FROM environment
        LEFT JOIN policy ON (policy.resource = 'environments/'||environment.resource_id AND policy.type = 'bb.policy.environment-tier')
        WHERE environment.deleted IS FALSE
        ORDER BY environment."order" ASC
    ) AS environments
ON CONFLICT (name) DO UPDATE SET value = EXCLUDED.value;


DROP TABLE IF EXISTS environment;
