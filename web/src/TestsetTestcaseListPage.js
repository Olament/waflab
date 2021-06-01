// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import React from "react";
import { Button, Col, Progress, Row, Switch, Table, Tag } from 'antd';
import * as Setting from "./Setting";
import { getStatusTagColor } from "./Setting";
import * as TestsetBackend from "./backend/TestsetBackend";
import * as TestcaseBackend from "./backend/TestcaseBackend";
import * as ResultBackend from "./backend/ResultBackend";

const ruleFileMap = {
  "905": "REQUEST-905-COMMON-EXCEPTIONS",
  "910": "REQUEST-910-IP-REPUTATION",
  "911": "REQUEST-911-METHOD-ENFORCEMENT",
  "912": "REQUEST-912-DOS-PROTECTION",
  "913": "REQUEST-913-SCANNER-DETECTION",
  "920": "REQUEST-920-PROTOCOL-ENFORCEMENT",
  "921": "REQUEST-921-PROTOCOL-ATTACK",
  "930": "REQUEST-930-APPLICATION-ATTACK-LFI",
  "931": "REQUEST-931-APPLICATION-ATTACK-RFI",
  "932": "REQUEST-932-APPLICATION-ATTACK-RCE",
  "933": "REQUEST-933-APPLICATION-ATTACK-PHP",
  "934": "REQUEST-934-APPLICATION-ATTACK-NODEJS",
  "941": "REQUEST-941-APPLICATION-ATTACK-XSS",
  "942": "REQUEST-942-APPLICATION-ATTACK-SQLI",
  "943": "REQUEST-943-APPLICATION-ATTACK-SESSION-FIXATION",
  "944": "REQUEST-944-APPLICATION-ATTACK-JAVA",
  "949": "REQUEST-949-BLOCKING-EVALUATION",
  "950": "RESPONSE-950-DATA-LEAKAGES",
  "951": "RESPONSE-951-DATA-LEAKAGES-SQL",
  "952": "RESPONSE-952-DATA-LEAKAGES-JAVA",
  "953": "RESPONSE-953-DATA-LEAKAGES-PHP",
  "954": "RESPONSE-954-DATA-LEAKAGES-IIS",
  "959": "RESPONSE-959-BLOCKING-EVALUATION",
  "980": "RESPONSE-980-CORRELATION",
};

class TestsetTestcaseListPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      testsetName: props.match.params.testsetName,
      testset: null,
      testcases: [],
      selectedRowKeys: [],
      isTarget: true,
    };
  }

  componentDidMount() {
    this.getTestset();
    this.getTestcases();
  }

  getTestset() {
    TestsetBackend.getTestset(this.state.testsetName)
      .then((testset) => {
        this.setState({
          testset: testset,
        });
      });
  }

  getTestcases() {
    TestcaseBackend.getFilteredTestcases(this.state.testsetName)
      .then((res) => {
        this.setState({
          testcases: res,
        });
      });
  }

  setTestcaseValue(i, key, value) {
    let testcases = this.state.testcases;
    testcases[i][key] = value;
    this.setState({
      testcases: testcases,
    });
  }

  getResult(testcase, i) {
    this.setTestcaseValue(i, "progressState", "ongoing");
    ResultBackend.getResult(this.state.testset.name, testcase.name, this.state.isTarget ? "target" : "baseline")
      .then((result) => {
        // Setting.showMessage("success", "Result: " + result.status);
        this.setTestcaseValue(i, "progressState", "finished");
        this.setTestcaseValue(i, "trueStatuses", result.statuses);
        this.setTestcaseValue(i, "response", result.response);
      })
      .catch(error => {
        Setting.showMessage("error", `failed to run: ${error}`);
        this.setTestcaseValue(i, "progressState", "error");
      });
  }

  getResults() {
    if (this.state.selectedRowKeys.length === 0) {
      this.state.testcases.forEach((testcase, i) => {
        this.getResult(testcase, i);
      });
    } else {
      this.state.testcases.forEach((testcase, i) => {
        if (this.state.selectedRowKeys.includes(testcase.name)) {
          this.getResult(testcase, i);
        }
      });
    }
  }

  getRuleId(record) {
    const start = record.name.indexOf("-");
    const end = record.name.indexOf(".");
    return record.name.substring(start + 1, end);
  }

  parseHitRules(record, hitRules) {
    const ruleId = this.getRuleId(record);

    if (hitRules === null) {
      hitRules = [];
    }
    const res = hitRules.map(hitRule => {
      const tokens = hitRule.split(",")
        .map(token => {
          let res = token.trim(" ");
          const i = res.lastIndexOf("-");
          res = res.substring(i + 1);
          return res;
        })
        .filter(token => {
          // remove hit rule token within INITIALIZATION range
          let value = parseInt(token, 10)
          return isNaN(value) || value > 901999
        });
      if (hitRule === "") {
        hitRule = "(empty)";
      } else {
        hitRule = tokens.join(", ");
      }

      let color;
      if (hitRule === "(empty)") {
        color = "400";
      } else if (hitRule.includes(ruleId)) {
        color = "200";
      } else {
        color = "403";
      }

      return { hitRule: hitRule, color: color };
    })
    return res;
  }

  renderTable(testcases) {
    const columns = [
      {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
        width: '150px',
        sorter: (a, b) => a.name.localeCompare(b.name),
        render: (text, record, index) => {
          return (
            <a target="_blank" href={`/testcases/${text}`}>{text}</a>
          )
        }
      },
      // {
      //   title: 'Description',
      //   dataIndex: 'desc',
      //   key: 'desc',
      //   width: '250px',
      //   sorter: (a, b) => a.desc.localeCompare(b.desc),
      // },
      {
        title: 'Rule',
        dataIndex: 'rule',
        key: 'rule',
        width: '80px',
        render: (text, record, index) => {
          const ruleId = this.getRuleId(record);
          const ruleFile = ruleFileMap[ruleId.substring(0, 3)];

          return (
            <Button onClick={() => Setting.openLink(`/rulesets/crs-3.2/rulefiles/${ruleFile}/rules/`)}>
              {ruleId}
            </Button>
          )
        }
      },
      // {
      //   title: 'Enabled',
      //   dataIndex: 'enabled',
      //   key: 'enabled',
      //   width: '80px',
      //   render: (text, record, index) => {
      //     return (
      //       <Switch disabled checked={text} />
      //     )
      //   }
      // },
      {
        title: '#Test',
        dataIndex: 'testCount',
        key: 'testCount',
        width: '100px',
        sorter: (a, b) => a.testCount - b.testCount,
      },
      // {
      //   title: 'Created Time',
      //   dataIndex: 'createdTime',
      //   key: 'createdTime',
      //   width: '160px',
      //   sorter: (a, b) => a.createdTime.localeCompare(b.createdTime),
      //   render: (text, record, index) => {
      //     return Setting.getFormattedDate(text);
      //   }
      // },
      {
        title: 'Method',
        dataIndex: 'method',
        key: 'method',
        width: '100px',
        sorter: (a, b) => a.method.localeCompare(b.method),
        render: (text, record, index) => {
          return Setting.getMethodTag(text);
        }
      },
      // {
      //   title: 'User-Agent',
      //   dataIndex: 'userAgent',
      //   key: 'userAgent',
      //   // width: '100px',
      //   sorter: (a, b) => a.userAgent.localeCompare(b.userAgent),
      // },
      {
        title: 'Expected',
        dataIndex: 'statusLists',
        key: 'statusLists',
        width: '600px',
        // ellipsis: true,
        // sorter: (a, b) => a.statusLists - b.statusLists,
        render: (text, record, index) => {
          return Setting.getStatusTags(text);
        }
      },
      {
        title: 'Baseline',
        dataIndex: 'baselineStatuses',
        key: 'baselineStatuses',
        width: '600px',
        // ellipsis: true,
        // sorter: (a, b) => a.trueStatuses - b.trueStatuses,
        render: (text, record, index) => {
          return Setting.getStatusTags(text);
        }
      },
      {
        title: 'Target',
        dataIndex: 'trueStatuses',
        key: 'trueStatuses',
        width: '600px',
        // ellipsis: true,
        // sorter: (a, b) => a.trueStatuses - b.trueStatuses,
        render: (text, record, index) => {
          return Setting.getStatusTags(text);
        }
      },
      {
        title: 'Result',
        dataIndex: 'result',
        key: 'result',
        width: '100px',
        sorter: (a, b) => a.result.localeCompare(b.result),
        render: (text, record, index) => {
          if (record.trueStatus > 0 && text === "") {
            return "(Empty)";
          } else {
            return text;
          }
        }
      },
      {
        title: 'Baseline Hit Rules',
        dataIndex: 'baseHitRules',
        key: 'baseHitRules',
        width: '100px',
        render: (text, record, index) => {
          if (text === null) {
            return "(Empty)";
          }

          const objs = this.parseHitRules(record, record.baselineHitRules);
          return objs.map(obj => {
            return (
              <Tag color={getStatusTagColor(obj.color)}>
                {Setting.getHighlightedText(obj.hitRule, this.getRuleId(record))}
              </Tag>
            );
          })
        }
      },
      {
        title: 'Hit Rules',
        dataIndex: 'hitRules',
        key: 'hitRules',
        width: '100px',
        render: (text, record, index) => {
          if (text === null) {
            return "(Empty)";
          }

          const objs = this.parseHitRules(record, record.hitRules);
          return objs.map(obj => {
            return (
              <Tag color={getStatusTagColor(obj.color)}>
                {Setting.getHighlightedText(obj.hitRule, this.getRuleId(record))}
              </Tag>
            );
          })
        }
      },
      {
        title: 'Default Action',
        dataIndex: 'action',
        key: 'action',
        width: '100px',
        sorter: (a, b) => a.action.localeCompare(b.action),
      },
      {
        title: 'Default State',
        dataIndex: 'state',
        key: 'state',
        width: '100px',
        sorter: (a, b) => a.state.localeCompare(b.state),
      },
      {
        title: 'Progress',
        key: 'progress',
        width: '100px',
        // sorter: (a, b) => a.userAgent.localeCompare(b.userAgent),
        render: (text, record, index) => {
          if (record.progressState === undefined) {
            return (
              <Progress percent={0} size="small" />
            )
          } else if (record.progressState === "ongoing") {
            return (
              <Progress percent={50} size="small" />
            )
          } else if (record.progressState === "error") {
            return (
              <Progress percent={100} size="small" status="exception" />
            )
          } else {
            return (
              <Progress percent={100} size="small" />
            )
          }
        }
      },
      {
        title: 'Action',
        dataIndex: '',
        key: 'op',
        width: '100px',
        render: (text, record, index) => {
          return (
            <div>
              <Button style={{ marginTop: '10px', marginBottom: '10px', marginRight: '10px' }}
                loading={record.progressState === "ongoing"} type="primary" onClick={() => this.getResult(record, index)}>Run</Button>
            </div>
          )
        }
      },
    ];

    const onSelectChange = selectedRowKeys => {
      this.setState({ selectedRowKeys });
    };

    const { selectedRowKeys } = this.state;
    const rowSelection = {
      selectedRowKeys,
      onChange: onSelectChange,
    };

    return (
      <div>
        <Table rowSelection={rowSelection} columns={columns} dataSource={testcases} rowKey="name" size="middle" bordered pagination={{ pageSize: 1000 }}
          title={() => (
            <div>
              <Tag color="#108ee9">{this.state.testset === null ? "" : this.state.testset.name}</Tag> Testcases&nbsp;&nbsp;&nbsp;&nbsp;
              <Button type="primary" size="small" onClick={this.getResults.bind(this)}>Run {this.state.selectedRowKeys.length === 0 ? "All" : "Selected"}</Button>&nbsp;&nbsp;&nbsp;&nbsp;
                   Target or Baseline ?&nbsp;&nbsp;
              <Switch checked={this.state.isTarget} onChange={(checked, e) => {
                this.setState({
                  isTarget: checked,
                });
              }} />
            </div>
          )}
          rowClassName={(record, index) => {
            if (record.action === "Block" && record.state === "Enabled" && record.result !== "ok: ") {
              return "red-row";
            } else if (record.action === "AnomalyScoring" && record.state === "Enabled" && this.parseHitRules(record, record.hitRules).some(pair => pair.color === "403")) {
              return "red-row";
            } else {
              return null;
            }
          }}
        />
      </div>
    );
  }

  render() {
    return (
      <div>
        <Row style={{ width: "100%" }}>
          {/*<Col span={1}>*/}
          {/*</Col>*/}
          <Col span={24}>
            {
              this.renderTable(this.state.testcases)
            }
          </Col>
          {/*<Col span={1}>*/}
          {/*</Col>*/}
        </Row>
      </div>
    );
  }
}

export default TestsetTestcaseListPage;
