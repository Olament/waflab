// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import React from "react";
import { Button, Col, List, Row, Table, Tag, Typography } from "antd";
import * as Setting from "./Setting";

const { Text } = Typography;

class RuleListPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      rulesetId: props.match.params.rulesetId,
      rulefileId: props.match.params.rulefileId,
      rulefile: null,
    };
  }

  componentDidMount() {
    this.listRules();
  }

  listRules() {
    fetch(`${Setting.ServerUrl}/api/list-rules?rulesetId=${this.state.rulesetId}&rulefileId=${this.state.rulefileId}`, {
      method: "GET",
      credentials: "include"
    })
      .then(res => res.json())
      .then((res) => {
        this.setState({
          rulefile: res,
        });
      }
      );
  }

  renderTable(title, title2, rules) {
    const columns = [
      // {
      //   title: 'No',
      //   dataIndex: 'no',
      //   key: 'no',
      //   width: 60,
      //   render: (text, record, index) => {
      //     return index;
      //   }
      // },
      {
        title: 'Id',
        dataIndex: 'id',
        key: 'id',
        width: 80,
      },
      // {
      //   title: 'Type',
      //   dataIndex: 'type',
      //   key: 'type',
      //   width: 80,
      // },
      {
        title: 'Paranoia Level',
        dataIndex: 'paranoiaLevel',
        key: 'paranoiaLevel',
        width: 80,
      },
      {
        title: 'Text',
        dataIndex: 'text',
        key: 'text',
      },
      {
        title: 'Tests',
        key: 'tests',
        width: 100,
        render: (text, record, index) => {
          if (record.regressionTestCount !== 0) {
            return (
              <div>
                <Button style={{ marginTop: '10px', marginBottom: '10px', marginRight: '10px' }} type="normal" onClick={() => Setting.openLink(`/testcases/${record.id}.yaml`)}>
                  {
                    `${record.regressionTestCount} Regression Tests`
                  }
                </Button>
              </div>
            )
          } else {
            return null;
          }
        }
      },
    ];

    const plColors = ["pl1", "pl2", "pl3", "pl4"];

    return (
      <div>
        <Table columns={columns} dataSource={rules} rowKey="text" size="middle" bordered pagination={{ pageSize: 1000 }}
          expandable={{
            defaultExpandAllRows: true,
            expandedRowRender: record => {
              return (
                <Row style={{ width: "100%" }}>
                  <Col span={6}>
                  </Col>
                  <Col span={18}>
                    <List
                      // bordered
                      dataSource={record.chainRules}
                      renderItem={(record, index) => (
                        <List.Item>
                          <Typography.Text mark>{`Chain rule - ${index}`}</Typography.Text>
                          {record.text}
                        </List.Item>
                      )}
                    />
                  </Col>
                </Row>
              )
            },
            rowExpandable: record => record.chainRules !== null,
          }}
          title={() => <div><Text>Rules for: </Text><Tag color="#108ee9">{title}</Tag> => <Tag color="#108ee9">{title2}</Tag></div>}
          rowClassName={(record, index) => { return plColors[record.paranoiaLevel - 1] }}
          loading={rules === null}
        />
      </div>
    );
  }

  render() {
    return (
      <div>
        <Row style={{ width: "100%" }}>
          <Col span={1}>
          </Col>
          <Col span={22}>
            {
              this.state.rulefile !== null ? this.renderTable(this.state.rulesetId, this.state.rulefileId, this.state.rulefile.rules) : null
            }
          </Col>
          <Col span={1}>
          </Col>
        </Row>
      </div>
    );
  }
}

export default RuleListPage;
