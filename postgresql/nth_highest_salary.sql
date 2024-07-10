CREATE OR REPLACE FUNCTION NthHighestSalary(N INT) RETURNS TABLE (Salary INT) AS $$
  BEGIN
    RETURN QUERY (
      SELECT CASE WHEN N > 0 THEN (
        SELECT DISTINCT e.salary
          FROM employee e
         ORDER BY e.salary DESC
        OFFSET N - 1
         LIMIT 1
      ) END
    );
  END;
$$ LANGUAGE plpgsql;
