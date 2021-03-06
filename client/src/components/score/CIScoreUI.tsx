import * as React from 'react';
import { Commit } from '../../types/commit';
import { CIScore } from 'src/types/score';
import { PieChart } from 'react-minimal-pie-chart';

export const CIScoreUI = (props: Props): React.ReactElement => {
  const [ciScore, setCiScore] = React.useState<CIScore | null>(null);
  React.useEffect(() => {
    if (ciScore == null) {
      (async () => {
        const result = await fetch(
          `http://localhost:8080/score/${props.commit.owner}/${props.commit.name}/ci?sha=${props.commit.commit?.end_commit.sha}`,
        );
        const json = await result.json();
        setCiScore(json);
      })();
    }
  }, [ciScore]);
  const color = ciScore ? (ciScore.score == 100 ? '#05c107' : ciScore.score > 50 ? '#E38627' : 'red') : 'black';
  return ciScore != null ? (
    <div style={{ textAlign: 'center' }}>
      <PieChart
        data={[{ value: ciScore.score, color: color }]}
        reveal={ciScore.score}
        lineWidth={20}
        background="#bfbfbf"
        animate
        label={() => ciScore.score}
        labelStyle={() => ({
          fill: color,
          fontSize: '25px',
          fontFamily: 'sans-serif',
        })}
        style={{ height: '150px' }}
        labelPosition={0}
      />
      {ciScore.score == 0 ? <p>Not Setting CI</p> : null}
      {ciScore.github_ci ? <p>Using GitHub CI</p> : null}
      {ciScore.travis_ci ? <p>Using Travis CI</p> : null}
      {ciScore.circle_ci ? <p>Using Circle CI</p> : null}
    </div>
  ) : (
    <div style={{ textAlign: 'center' }}>loading</div>
  );
};

type Props = {
  commit: Commit;
};
